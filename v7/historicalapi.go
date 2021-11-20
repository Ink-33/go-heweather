package v7

import "time"

// NewHistoricalClient 创建一个历史数据查询实例。
// historicalType支持传入weather，air
// https://dev.heweather.com/docs/api/historical
func NewHistoricalClient(location, date, historicalType string) (client HeWeatherAPI, err error) {
	if historicalType != "weather" && historicalType != "air" {
		err = &NewClientErr{Reason: "Invalid historicalType: " + historicalType}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "historical",
		Parameter: map[string]string{"date": date, "location": location},
		SubName:   historicalType,
		Timeout:   15 * time.Second,
	}, nil
}
