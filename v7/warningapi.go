package v7

import "time"

// NewWarningClient 创建一个灾害预警查询实例。
// https://dev.qweather.com/docs/api/warning/weather-warning/
func NewWarningClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "warning",
		Parameter: map[string]string{"location": location},
		SubName:   "now",
		Timeout:   15 * time.Second,
	}
}

// NewWarningListClient 创建一个灾害预警城市列表查询实例。
// 当前WarningRange仅支持cn
// https://dev.qweather.com/docs/api/warning/weather-warning-city-list/
func NewWarningListClient(warningRange string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "warning",
		Parameter: map[string]string{"range": warningRange},
		SubName:   "list",
		Timeout:   15 * time.Second,
	}
}
