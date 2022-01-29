package v7

import "time"

// NewLiveIndexClient 创建一个生活指数查询实例。
// https://dev.heweather.com/docs/api/indices
func NewLiveIndexClient(location, indexType, duration string) (client HeWeatherAPI, err error) {
	if duration != "1d" && duration != "3d" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "indices",
		Parameter: map[string]string{"location": location, "type": indexType},
		SubName:   duration,
		Timeout:   15 * time.Second,
	}, nil
}
