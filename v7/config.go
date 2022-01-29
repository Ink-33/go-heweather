package v7

import (
	"time"
)

// CNAPIAddress 和风天气商业版中国节点地址，当某些API在海外访问返回401时请尝试使用此地址
const CNAPIAddress = "https://cn-api.qweather.com/v7"

// SetAPIOptionParam 设置API可选参数
func (u *universeHeWeatherAPI) SetAPIOptionParam(config map[string]string) {
	if len(config) == 0 {
		return
	}
	for k, v := range config {
		if u.Parameter[k] == "" {
			u.Parameter[k] = v
		}
	}

}

// SetTimeout 设置超时时间
func (u *universeHeWeatherAPI) SetTimeout(timeout time.Duration) {
	u.Timeout = timeout
}

// SetCustomAPIAddress 设置自定义API地址
func (u *universeHeWeatherAPI) SetCustomAPIAddress(address string) {
	u.CustomAPIAddress = address
}
