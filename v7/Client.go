package v7

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
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

// APIConfig 用于配置天气API的各种配置
type APIConfig struct {
	p       map[string]string // 参数
	Timeout time.Duration     // 超时时间
}

// NewAPIConfig 创建一个API配置
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		p: make(map[string]string),
	}
}

// SetLanguage 设置语言
func (c *APIConfig) SetLanguage(lang string) {
	if c.p == nil {
		c.p = make(map[string]string)
	}
	c.p["lang"] = lang
}

// SetUnit 设置单位
func (c *APIConfig) SetUnit(unit string) {
	if c.p == nil {
		c.p = make(map[string]string)
	}
	c.p["unit"] = unit
}

// SetAdm 设置行政区划
func (c *APIConfig) SetAdm(adm string) {
	if c.p == nil {
		c.p = make(map[string]string)
	}
	c.p["adm"] = adm
}

// SetRange 设置范围
func (c *APIConfig) SetRange(rangeStr string) {
	if c.p == nil {
		c.p = make(map[string]string)
	}
	c.p["range"] = rangeStr
}

// SetNumber 设置数量
func (c *APIConfig) SetNumber(number string) {
	if c.p == nil {
		c.p = make(map[string]string)
	}
	c.p["number"] = number
}

// SetTimeout 设置超时时间
func (c *APIConfig) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
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
	map1 := u.APIConfig.p
	var map2 = make(map[string]string)
	for k, v := range map1 {
		if map2[k] == "" {
			map2[k] = v
		}
	}
	for k, v := range u.Parameter {
		if map2[k] == "" {
			map2[k] = v
		}
	}
	paramstr, signature := GetSignature(credential.PublicID, credential.Key, map2)
	result, err = httpClient(urlBuilder(u.getURL(credential), u.Name, u.SubName)+"?"+paramstr+"&sign="+signature, u.APIConfig.Timeout)
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
	req.Proto = "HTTP/1.1"
	req.Header.Add("User-Agent", "go-heweather SDK")
	rep, err := httpc.Do(req)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()
	content, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
