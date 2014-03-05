Accuweather Api url generate
========

##描述

使用accuweather生成访问url。

##安装方法

	cd $GOPATH/src
	git clone https://github.com/nosix-me/accuweather.git github.com/nosix-me/accuweather
	go install github.com/nosix-me/accuweather

##使用方法

>引入

	import "github.com/nosix-me/accuweather"
	import "github.com/nosix-me/accuweather/params"

>初始化

	accu, err := accuweather.SetKeys("apiKey","locationKey")

>获得请求的url
	
	url,err := accu.GetCurrentConditionsUrl(params.ENVRIOMENT_DEVELOPMENT, params.VERSION_1, params.FORMAT_JSON, params.LANGUAGE_CHINESE, params.DETAILS_TRUE)

---


	
