package v7

import "time"

// NewTideClient 创建一个潮汐查询实例。
// https://dev.qweather.com/docs/api/ocean/tide/
func NewOceanTideClient(location, date string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "ocean",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   "tide",
		Timeout:   15 * time.Second,
	}
}

// NewOceanCurrentsClient 创建一个潮流查询实例。
// https://dev.qweather.com/docs/api/ocean/currents/
func NewOceanCurrentsClient(location, date string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "ocean",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   "currents",
		Timeout:   15 * time.Second,
	}
}
