package v7

import "time"

// NewAirQualityClient 创建一个空气质量查询实例。
// duration仅支持 now,5d
// https://dev.heweather.com/docs/api/air
func NewAirQualityClient(location, duration string) (client HeWeatherAPI, err error) {
	if duration != "5d" && duration != "now" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "air",
		Parameter: map[string]string{"location": location},
		SubName:   duration,
		Timeout:   15 * time.Second,
	}, nil
}
