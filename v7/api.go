package v7

import "time"

// HeWeatherAPI 和风天气v7API通用接口
type HeWeatherAPI interface {
	// Run 执行API
	Run(credential *Credential) (result string, err error)
	// GetURL 获取API链接
	getURL(credential *Credential) (url string)
	// SetAPIOptionParam 设置API可选参数
	SetAPIOptionParam(config map[string]string)
	// SetTimeout 设置超时时间
	SetTimeout(timeout time.Duration)
}

type universeHeWeatherAPI struct {
	isGeo     bool
	Name      string
	Parameter map[string]string
	SubName   string
	Timeout   time.Duration
}

// NewClientErr 创建查询实例时返回的错误
type NewClientErr struct {
	Reason string
}

func (e *NewClientErr) Error() string {
	return e.Reason
}

// NewGeoCityClient 创建一个城市信息搜索实例。
// https://dev.heweather.com/docs/api/geo#%E5%9F%8E%E5%B8%82%E4%BF%A1%E6%81%AF%E6%90%9C%E7%B4%A2
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
// https://dev.heweather.com/docs/api/geo#%E7%83%AD%E9%97%A8%E5%9F%8E%E5%B8%82%E6%9F%A5%E8%AF%A2
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
// 当前POItype仅限景点。例如scenic
// https://dev.heweather.com/docs/api/geo#poi%E4%BF%A1%E6%81%AF%E6%90%9C%E7%B4%A2
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
// 当前POItype仅限景点。例如scenic
// https://dev.heweather.com/docs/api/geo#poi%E8%8C%83%E5%9B%B4%E6%90%9C%E7%B4%A2
func NewGeoPOIRangeClient(location string, poiType string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     true,
		Name:      "poi",
		Parameter: map[string]string{"location": location, "type": poiType},
		SubName:   "range",
		Timeout:   15 * time.Second,
	}
}

// NewRealTimeWeatherClient 创建一个实况天气查询实例。
// https://dev.heweather.com/docs/api/weather
func NewRealTimeWeatherClient(location string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "weather",
		Parameter: map[string]string{"location": location},
		SubName:   "now",
		Timeout:   15 * time.Second,
	}
}

// NewWeatherForecastClient 创建一个天气预报查询实例。
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

// NewWarningClient 创建一个灾害预警查询实例。
// https://dev.heweather.com/docs/api/warning
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
// https://dev.qeweather.com/docs/api/warning
func NewWarningListClient(warningRange string) (client HeWeatherAPI) {
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "warning",
		Parameter: map[string]string{"range": warningRange},
		SubName:   "list",
		Timeout:   15 * time.Second,
	}
}

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

// NewWeatherPOIClient 创建一个景区天气查询实例。
// https://dev.heweather.com/docs/api/weather-poi
func NewWeatherPOIClient(location, duration string) (client HeWeatherAPI, err error) {
	if duration != "7d" && duration != "now" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	return &universeHeWeatherAPI{
		isGeo:     false,
		Name:      "weather-poi",
		Parameter: map[string]string{"location": location},
		SubName:   duration,
		Timeout:   15 * time.Second,
	}, nil
}

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

// NewSunandMoonClient 创建一个日出日落、月升月落和月相查询实例。
// historicalType支持传入weather，air
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
