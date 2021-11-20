package v7

import "time"

// NewSunandMoonClient 创建一个日出日落、月升月落和月相查询实例。
// https://dev.heweather.com/docs/api/astronomy
func NewSunandMoonClient(location, date string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "astronomy",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   "sunmoon",
		Timeout:   15 * time.Second,
	}
}

// NewSunClient 创建一个日出日落查询实例。
// https://dev.qweather.com/docs/api/astronomy/sunrise-sunset/
func NewSunClient(location, date string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "astronomy",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   "sun",
		Timeout:   15 * time.Second,
	}
}

// NewmoonClient 创建一个月升月落和月相查询实例。
// https://dev.qweather.com/docs/api/astronomy/moon-and-moon-phase/
func NewMoonClient(location, date string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "astronomy",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   "moon",
		Timeout:   15 * time.Second,
	}
}

// NewSolarElevationAngleClient 创建一个太阳高度查询实例。
// https://dev.qweather.com/docs/api/astronomy/solar-elevation-angle/
func NewSolarElevationAngleClient(location, date, timeStr, tz, alt string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "astronomy",
		Parameter: map[string]string{"date": date, "location": location, "time": timeStr, "tz": tz, "alt": alt},
		SubName:   "solar_elevation_angle",
		Timeout:   15 * time.Second,
	}
}
