package v7

import "time"

// HeWeatherAPI 和风天气v7API通用接口
type HeWeatherAPI interface {
	// Run 执行API
	Run(credential *Credential) (result string, err error)
	// GetURL 获取API链接
	getURL(credential *Credential) (url string)
	// SetAPIOptionParam 设置API可选参数
	SetAPIOptionParam(config map[string]string)
	// SetTimeout 设置超时时间
	SetTimeout(timeout time.Duration)
}

type universeHeWeatherAPI struct {
	isGeo     bool
	Name      string
	Parameter map[string]string
	SubName   string
	Timeout   time.Duration
}

// NewClientErr 创建查询实例时返回的错误
type NewClientErr struct {
	Reason string
}

func (e *NewClientErr) Error() string {
	return e.Reason
}


// Deprecated: 该API已被移除
// NewWeatherPOIClient 创建一个景区天气查询实例。
// https://dev.heweather.com/docs/api/weather-poi
func NewWeatherPOIClient(location, duration string) (client HeWeatherAPI, err error) {
	if duration != "7d" && duration != "now" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "weather-poi",
		Parameter: map[string]string{"location": location},
		SubName:   duration,
		Timeout:   15 * time.Second,
	}, nil
}

