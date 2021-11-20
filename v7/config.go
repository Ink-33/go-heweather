package v7

import (
	"time"
)

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

