package v7

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Credential 和风天气凭证
type Credential struct {
	PublicID   string
	Key        string
	IsBusiness bool
}

// NewCredential 创建一个和风天气凭证
func NewCredential(publicID, key string, isBusiness bool) (credential *Credential) {
	credential = &Credential{
		PublicID:   publicID,
		Key:        key,
		IsBusiness: isBusiness,
	}
	return
}

func (u *universeHeWeatherAPI) Run(credential *Credential) (result string, err error) {
	paramstr, signature := GetSignature(credential.PublicID, credential.Key, u.Parameter)
	result, err = httpClient(urlBuilder(u.getURL(credential), u.Name, u.SubName)+"?"+paramstr+"&sign="+signature, u.Timeout)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (u *universeHeWeatherAPI) getURL(credential *Credential) (url string) {
	if u.isGeo {
		return "https://geoapi.qweather.net/v2/"
	}
	if credential.IsBusiness {
		return "https://api.qweather.com/v7"
	}
	return "https://devapi.qweather.com/v7"
}

func urlBuilder(url, name, subName string) string {
	return fmt.Sprintf("%s/%s/%s", url, name, subName)
}

// GetSignature 和风天气签名生成算法-Golang版本
func GetSignature(publicID, key string, param map[string]string) (paramstr, signature string) {
	sa := []string{}
	for k, v := range param {
		if v != "" {
			sa = append(sa, k+"="+v)
		}
	}
	sa = append(sa, "t="+strconv.FormatInt(time.Now().Unix(), 10))
	sa = append(sa, "username="+publicID)
	sort.Strings(sa)
	paramstr = strings.Join(sa, "&")
	md5c := md5.New()
	md5c.Reset()
	_, _ = md5c.Write([]byte(paramstr + key))
	return paramstr, fmt.Sprintf("%x", md5c.Sum(nil))
}

func httpClient(address string, timeout time.Duration) (result string, err error) {
	httpc := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "go-heweather SDK")
	rep, err := httpc.Do(req)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()
	content, err := io.ReadAll(rep.Body)
	println(rep.Proto, rep.Status)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
