package v7

//HeWeatherAPI 和风天气v7API通用接口
type HeWeatherAPI interface {
	//Run 执行API
	Run(credential *Credential, config *ClientConfig) (Result string, err error)
	//GetURL 获取API链接
	GetURL(credential *Credential) (URL string)
}

//HeWeaGeoAPI 和风天气城市信息搜索接口
type HeWeaGeoAPI interface {
	//Run 执行API
	Run(credential *Credential, config *ClientConfig) (Result string, err error)
	//GetURL 获取API链接
	GetURL() (URL string)
}

type universeHeWeatherAPI struct {
	Name      string
	SubName   string
	Parameter map[string]string
}
type geoAPI struct {
	Name    string
	SubName string
	Locaton string
}

//NewClientErr 创建查询实例时返回的错误
type NewClientErr struct {
	Reason string
}

func (e *NewClientErr) Error() string {
	return e.Reason
}

//NewGeoClient 创建一个城市信息搜索实例
//https://dev.heweather.com/docs/api/geo
func NewGeoClient(location string) (Client HeWeaGeoAPI) {
	Client = &geoAPI{
		Name:    "city",
		SubName: "lookup",
		Locaton: location,
	}
	return
}

//NewRealTimeWeatherClient 创建一个实况天气查询实例。
//https://dev.heweather.com/docs/api/weather
func NewRealTimeWeatherClient(location string) (Client HeWeatherAPI) {
	p := map[string]string{"location": location}
	Client = &universeHeWeatherAPI{
		Name:      "weather",
		SubName:   "now",
		Parameter: p,
	}
	return
}

//NewWeatherForecastClient 创建一个天气预报查询实例。
//你需要在 https://dev.heweather.com/docs/api/weather 查询指定的Duration，
//如3d，24h
func NewWeatherForecastClient(location string, duration string) (Client HeWeatherAPI, err error) {
	d := []string{"3d", "7d", "10d", "15d", "24h", "72h", "168h"}
	p := map[string]string{"location": location}
	var dr string = ""
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
	Client = &universeHeWeatherAPI{
		Name:      "weather",
		SubName:   dr,
		Parameter: p,
	}
	return
}

//NewMinutelyClient 创建一个分钟级降水查询实例。
//https://dev.heweather.com/docs/api/minutely
func NewMinutelyClient(location string) (Client HeWeatherAPI) {
	p := map[string]string{"location": location}
	Client = &universeHeWeatherAPI{
		Name:      "minutely",
		SubName:   "5m",
		Parameter: p,
	}
	return
}

//NewAirQualityClient 创建一个空气质量查询实例。
//duration仅支持 now,5d
//https://dev.heweather.com/docs/api/air
func NewAirQualityClient(location, duration string) (Client HeWeatherAPI, err error) {
	p := map[string]string{"location": location}
	if duration != "5d" && duration != "now" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	Client = &universeHeWeatherAPI{
		Name:      "air",
		SubName:   duration,
		Parameter: p,
	}
	return
}

//NewWarningClient 创建一个灾害预警查询实例。
//https://dev.heweather.com/docs/api/warning
func NewWarningClient(location string) (Client HeWeatherAPI) {
	p := map[string]string{"location": location}
	Client = &universeHeWeatherAPI{
		Name:      "warning",
		SubName:   "now",
		Parameter: p,
	}
	return
}

//NewWarningListClient 创建一个灾害预警城市列表查询实例。
//当前WarningRange仅支持cn
//https://dev.heweather.com/docs/api/warning
func NewWarningListClient(WarningRange string) (Client HeWeatherAPI) {
	p := map[string]string{"range": WarningRange}
	Client = &universeHeWeatherAPI{
		Name:      "warning",
		SubName:   "now",
		Parameter: p,
	}
	return
}

//NewLiveIndexClient 创建一个生活指数查询实例。
//https://dev.heweather.com/docs/api/indices
func NewLiveIndexClient(location, indexType, duration string) (Client HeWeatherAPI, err error) {
	p := map[string]string{"location": location, "type": indexType}
	if duration != "1d" && duration != "3d" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	Client = &universeHeWeatherAPI{
		Name:      "indices",
		SubName:   duration,
		Parameter: p,
	}
	return
}

//NewWeatherPOIClient 创建一个景区天气查询实例。
//https://dev.heweather.com/docs/api/weather-poi
func NewWeatherPOIClient(location, duration string) (Client HeWeatherAPI, err error) {
	p := map[string]string{"location": location}
	if duration != "7d" && duration != "now" {
		err = &NewClientErr{Reason: "Invalid duration: " + duration}
		return nil, err
	}
	Client = &universeHeWeatherAPI{
		Name:      "weather-poi",
		SubName:   duration,
		Parameter: p,
	}
	return
}

//NewHistoricalClient 创建一个历史数据查询实例。
//historicalType支持传入weather，air
//https://dev.heweather.com/docs/api/historical
func NewHistoricalClient(location, date, historicalType string) (Client HeWeatherAPI, err error) {
	if historicalType != "weather" && historicalType != "air" {
		err = &NewClientErr{Reason: "Invalid historicalType: " + historicalType}
		return nil, err
	}
	p := map[string]string{"date": date, "location": location}
	Client = &universeHeWeatherAPI{
		Name:      "historical",
		SubName:   historicalType,
		Parameter: p,
	}
	return
}

//NewSunandMoonClient 创建一个日出日落、月升月落和月相查询实例。
//historicalType支持传入weather，air
//https://dev.heweather.com/docs/api/astronomy
func NewSunandMoonClient(location, date string) (Client HeWeatherAPI) {
	p := map[string]string{"date": date, "location": location}
	Client = &universeHeWeatherAPI{
		Name:      "astronomy",
		SubName:   "sunmoon",
		Parameter: p,
	}
	return
}
