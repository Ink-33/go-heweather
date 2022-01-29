package v7

import "time"

// basins 台风列表查询可用的basin参数(需要查询的台风所在的流域)
type basins struct {
	NorthAtlantic       string
	EasternPacific      string
	NorthWestPacific    string
	SouthWesternPacific string
	NorthIndian         string
	SouthIndian         string
}

// Basins 台风列表查询可用的basin参数(需要查询的台风所在的流域)
var Basins = basins{
	NorthAtlantic:       "AL",
	EasternPacific:      "EP",
	NorthWestPacific:    "NP",
	SouthWesternPacific: "SP",
	NorthIndian:         "NI",
	SouthIndian:         "SI",
}

// NewStormListClient 创建一个台风列表查询实例。
// https://dev.qweather.com/docs/api/tropical/storm-list/
func NewStormListClient(basin, year string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "tropical",
		Parameter: map[string]string{"basin": basin, "year": year},
		SubName:   "storm-list",
		Timeout:   15 * time.Second,
	}
}

// NewStormTrackClient 创建一个台风实况和路径查询实例。
// https://dev.qweather.com/docs/api/tropical/storm-track/
func NewStormTrackClient(stormID string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "tropical",
		Parameter: map[string]string{"stormid": stormID},
		SubName:   "storm-track",
		Timeout:   15 * time.Second,
	}
}

// NewStormForecastClient 创建一个台风预报查询实例。
// https://dev.qweather.com/docs/api/tropical/storm-forecast/
func NewStormForecastClient(stormID string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "tropical",
		Parameter: map[string]string{"stormid": stormID},
		SubName:   "storm-forecast",
		Timeout:   15 * time.Second,
	}
}
