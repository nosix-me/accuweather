accuweather api url generate
========
[![Build Status](https://drone.io/github.com/widuu/goini/status.png)](https://drone.io/github.com/nosix-me/go-properties/6)

##描述

使用accuweather生成访问url。

##安装方法
	
	go get github.com/nosix-me/accuweather

##使用方法

>引入

	import "github.com/nosix-me/accuweather"
	import "github.com/nosix-me/accuweather/params"

>初始化

	accu, err := accuweather.SetKeys("apiKey","locationKey")

>获得请求的url
	
	url,err := accu.GetCurrentConditionsUrl(params.ENVRIOMENT_DEVELOPMENT, params.VERSION_1, params.FORMAT_JSON, params.LANGUAGE_CHINESE, params.DETAILS_TRUE)

---


	
