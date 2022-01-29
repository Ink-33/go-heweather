package v7

import "time"

// NewGeoCityClient 创建一个城市信息搜索实例。
// https://dev.qweather.com/docs/api/geo/city-lookup/
func NewGeoCityClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     true,
		Name:      "city",
		Parameter: map[string]string{"location": location},
		SubName:   "lookup",
		Timeout:   15 * time.Second,
	}
}

// NewGeoTopCityClient 创建一个热门城市查询实例。
// https://dev.qweather.com/docs/api/geo/top-city/
func NewGeoTopCityClient() (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     true,
		Name:      "city",
		Parameter: map[string]string{},
		SubName:   "top",
		Timeout:   15 * time.Second,
	}
}

// NewGeoPOIClient 创建一个POI信息搜索实例。
// https://dev.qweather.com/docs/api/geo/poi-lookup/
func NewGeoPOIClient(location string, poiType string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     true,
		Name:      "poi",
		Parameter: map[string]string{"location": location, "type": poiType},
		SubName:   "lookup",
		Timeout:   15 * time.Second,
	}
}

// NewGeoPOIRangeClient 创建一个POI范围搜索实例。
// https://dev.qweather.com/docs/api/geo/poi-range/
func NewGeoPOIRangeClient(location string, poiType string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     true,
		Name:      "poi",
		Parameter: map[string]string{"location": location, "type": poiType},
		SubName:   "range",
		Timeout:   15 * time.Second,
	}
}
