package v7

import "time"

// NewRealTimeWeatherClient 创建一个实况天气查询实例。
// https://dev.qweather.com/docs/api/weather/weather-now/
func NewRealTimeWeatherClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "weather",
		Parameter: map[string]string{"location": location},
		SubName:   "now",
		Timeout:   15 * time.Second,
	}
}

// NewWeatherForecastClient 创建一个天气预报查询实例(包括逐天与逐小时)。
// 你需要在 https://dev.heweather.com/docs/api/weather 查询指定的Duration，
// 如3d，24h
func NewWeatherForecastClient(location string, duration string) (client HeWeatherAPI, err error) {
	d := []string{"3d", "7d", "10d", "15d", "24h", "72h", "168h"}
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
		Name:      "weather",
		Parameter: map[string]string{"location": location},
		SubName:   dr,
		Timeout:   15 * time.Second,
	}, nil
}
