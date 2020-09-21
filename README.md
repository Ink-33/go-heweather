# go-heweather
[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://pkg.go.dev/mod/github.com/Ink-33/go-heweather)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ink-33/go-heweather)](https://goreportcard.com/report/github.com/Ink-33/go-heweather) 

[和风天气](https://heweather.com/)的非官方Golang SDK

# 简介
欢迎使用go-heweather！  
本SDK基于和风天气WebAPI V7而制作，方便Go开发者快速调用和风天气API。  
您亦可在GoDoc[here](https://pkg.go.dev/mod/github.com/Ink-33/go-heweather) 上浏览本SDK的文档，在`package`中选择v7即可浏览和风天气v7的相关文档。

# 开始使用
使用本SDK，您需要一个和风天气WebAPI的安全凭证。可前往[和风天气开发平台](https://dev.heweather.com/)获取您的安全凭证。

## 获取SDK

    go get github.com/Ink-33/go-heweather

# 示例
和风天气v7的每个API都有其对应的一个函数，通过此函数可以获取一个请求实例。以下是使用例程。  
确保您使用本sdk前已阅读[和风天气开发文档](https://dev.heweather.com/docs/start/)

### 最简例程
以下为完成一次请求的所有必须操作
```go
package main

import hewea "github.com/Ink-33/go-heweather/v7"

func main() {
	var publicID = "your public ID"
	var key = "your key"
	//免费开发版为false，商业共享版与商业高性能版均为true
	var isBusiness = false
	//创建一个安全凭证
	credential := hewea.NewCredential(publicID, key, isBusiness)
	//要查询的地址
    var location = "101010100"
    //新建一个实时天气查询实例
	client := hewea.NewRealTimeWeatherClient(location)
	//运行
	rep, err := client.Run(credential, nil)
	if err != nil {
		panic(err) //也可以自行进行错误处理
	}
	println(rep)
}

```
要注意的是，当前版本的sdk不会对调用返回值进行进一步处理，而是直接返回接口的返回值，您应当根据您的需求进一步对返回值进行处理。

### 高级篇
您还可以通过以下方法对api进行定制
``` go
	var location = "101010100"
	//查询时间段
    var duration = "now"
    //新建一个空气质量查询实例
	client, err := hewea.NewAirQualityClient(location, duration)
	if err != nil {
		panic(err) //此处返回错误代表您填入了一个错误的duration
	}
	//此处新建一个请求实例配置。
	//请注意，您要请求的api并不一定支持全部配置，请您按需填写
	//各api支持的配置请参考https://dev.heweather.com/docs/api
	//本结构体支持的配置请参考GoDoc
	cpf := &hewea.ClientConfig{
		Language: "cn",
	}
	//运行
	rep, err := client.Run(credential, cpf)
	if err != nil {
		panic(err) //也可以自行进行错误处理
	}
	println(rep)
}

```

