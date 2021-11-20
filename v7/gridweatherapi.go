package v7

import "time"

// NewMinutelyClient 创建一个分钟级降水查询实例。
// https://dev.heweather.com/docs/api/minutely
func NewMinutelyClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "minutely",
		Parameter: map[string]string{"location": location},
		SubName:   "5m",
		Timeout:   15 * time.Second,
	}
}

// NewGridRealTimeWeatherClient 创建一个格点实况天气查询实例。
// https://dev.qweather.com/docs/api/grid-weather/grid-weather-now/
func NewGridRealTimeWeatherClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "grid-weather",
		Parameter: map[string]string{"location": location},
		SubName:   "now",
		Timeout:   15 * time.Second,
	}
}

// NewGridWeatherForecastClient 创建一个格点天气预报查询实例(包括逐天与逐小时)。
// 你需要在 https://dev.qweather.com/docs/api/grid-weather 查询指定的Duration，
// 如3d，24h
func NewGridWeatherForecastClient(location string, duration string) (client HeWeatherAPI, err error) {
	d := []string{"3d", "7d", "24h"}
	dr := ""
	for _, v := range d {
		if duration == v {
			dr = duration
			break
		}
	}
	if dr == "" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "grid-weather",
		Parameter: map[string]string{"location": location},
		SubName:   dr,
		Timeout:   15 * time.Second,
	}, nil
}
