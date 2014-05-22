// accuweather project accuweather.go
//auther:Yaxiang He
package accuweather

import (
	"errors"
)

type Accuweather struct {
	apiKey      string
	locationKey string
}

type Params string

//Set apiKey and locationKey
func SetKeys(apiKey, locationKey string) (*Accuweather, error) {
	if len(apiKey) == 0 || len(locationKey) == 0 {
		return nil, errors.New("ApiKey or LocationKey can't be empty!")
	}
	accu := new(Accuweather)
	accu.apiKey = apiKey
	accu.locationKey = locationKey

	return accu, nil
}

//Get current condition url
func (accu *Accuweather) GetCurrentConditionsUrl(enviroment, version, format, language, details Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys error!")
	}
	return string(enviroment) + string(TYPES_CURRENTCONDITIONS) + string(version) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get current time prev 6 hours history data
func (accu *Accuweather) GetHistoricalCurrentConditionsUrl(enviroment, version, format, language, details Params, locationKey, apiKey string) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys error!")
	}
	return string(enviroment) + string(TYPES_CURRENTCONDITIONS) + string(version) + accu.locationKey + "/historical" + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get future weather data for hourly
func (accu *Accuweather) GetForecastHourly(enviroment Params, forecasthourly Params, version Params, format Params, locationKey string, apiKey string, language Params, details Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys error!")
	}
	return string(enviroment) + string(TYPES_FORECASTS) + string(version) + string(FROECAST_HOURLY) + string(forecasthourly) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get future weather data for daily
func (accu *Accuweather) GetForecastDaily(enviroment Params, forecastdaily Params, version Params, format Params, locationKey string, apiKey string, language Params, details Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys error!")
	}
	return string(enviroment) + string(TYPES_FORECASTS) + string(version) + string(FROECAST_DAILY) + string(forecastdaily) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

const (

	// details

	DETAILS_TRUE  Params = "true"
	DETAILS_FALSE Params = "false"

	// accuweather development enviroment

	ENVRIOMENT_DEVELOPMENT Params = "http://apidev.accuweather.com/"
	ENVRIOMENT_PRODUCTION  Params = "http://api.accuweathet.com/"

	// forecasts daily

	DAILY_DAY_1  Params = "1day/"
	DAILY_DAY_5  Params = "5day/"
	DAILY_DAY_10 Params = "10day/"
	DAILY_DAY_15 Params = "15day/"
	DAILY_DAY_25 Params = "25day/"

	// forecast hourly

	HOURLY_HOUR_1   Params = "1hour/"
	HOURLY_HOUR_12  Params = "12hour/"
	HOURLY_HOUR_24  Params = "24hour/"
	HOURLY_HOUR_72  Params = "72hour/"
	HOURLY_HOUR_120 Params = "120hour/"
	HOURLY_HOUR_240 Params = "240hour/"

	// forecast

	FROECAST_HOURLY Params = "hourly/"
	FROECAST_DAILY  Params = "daily/"

	// format

	FORMAT_JSON Params = ".json?"

	// accuweather support's language

	LANGUAGE_ARABIC                 Params = "ar"
	LANGUAGE_ARABIC_ALGERIA         Params = "ar-dz"
	LANGUAGE_ARABIC_BAHRAIN         Params = "ar-bh"
	LANGUAGE_ARABIC_EGYPT           Params = "ar-eg"
	LANGUAGE_ARABIC_IRAQ            Params = "ar-iq"
	LANGUAGE_ARABIC_JORDAN          Params = "ar-jo"
	LANGUAGE_ARABIC_KUWAIT          Params = "ar-kw"
	LANGUAGE_ARABIC_LEBANON         Params = "ar-lb"
	LANGUAGE_ARABIC_LIBYA           Params = "ar-ly"
	LANGUAGE_ARABIC_MOROCCO         Params = "ar-ma"
	LANGUAGE_ARABIC_OMAN            Params = "ar-om"
	LANGUAGE_ARABIC_QATAR           Params = "ar-qa"
	LANGUAGE_ARABIC_SAUDI_ARABIA    Params = "ar-sa"
	LANGUAGE_ARABIC_SYRIA           Params = "ar-sy"
	LANGUAGE_ARABIC_TUNISIA         Params = "ar-tn"
	LANGUAGE_ARABIC_U_A_E           Params = "ar-ae"
	LANGUAGE_ARABIC_YEMEN           Params = "ar-ye"
	LANGUAGE_BULGARIAN              Params = "bg"
	LANGUAGE_CATALAN                Params = "ca"
	LANGUAGE_CROATIAN               Params = "hr"
	LANGUAGE_CZECH                  Params = "cs"
	LANGUAGE_CHINESE                Params = "zh"
	LANGUAGE_CHINESE_HK             Params = "zh-hk"
	LANGUAGE_CHINESE_PRC            Params = "zh-cn"
	LANGUAGE_CHINESE_SINGAPORE      Params = "zh-sg"
	LANGUAGE_CHINESE_TAIWAN         Params = "zh-tw"
	LANGUAGE_DANISH                 Params = "da"
	LANGUAGE_DUTCH                  Params = "nl"
	LANGUAGE_DUTCH_BELGIUM          Params = "nl-be"
	LANGUAGE_ENGLISH                Params = "en"
	LANGUAGE_ENGLISH_AUSTRALIA      Params = "en-au"
	LANGUAGE_ENGLISH_BELIZE         Params = "en-bz"
	LANGUAGE_ENGLISH_CANADA         Params = "en-ca"
	LANGUAGE_ENGLISH_IRELAND        Params = "en-ie"
	LANGUAGE_ENGLISH_NEW_ZEALANG    Params = "en-nz"
	LANGUAGE_ENGLIST_SOUTH_AFRICA   Params = "en-za"
	LANGUAGE_ENGLISH_TRINIDAD       Params = "en-tt"
	LANGUAGE_ENGLISH_UNITED_KINGDOM Params = "en-gb"
	LANGUAGE_ENGLISH_UNITED_STATES  Params = "en-us"
	LANGUAGE_ESTONIAN               Params = "et"
	LANGUAGE_FILIPINO               Params = "ph"
	LANGUAGE_FINNISH                Params = "fi"
	LANGUAGE_FRENCH                 Params = "fr"
	LANGUAGE_FRENCH_BELGIUM         Params = "fr-be"
	LANGUAGE_FRENCH_CANADA          Params = "fr-ca"
	LANGUAGE_FRENCH_LUXEMBOURG      Params = "fr-lu"
	LANGUAGE_FRENCH_SWITZERLAND     Params = "fr-ch"

	// accuweather request

	TYPES_CURRENTCONDITIONS Params = "currentconditions/"
	TYPES_FORECASTS         Params = "forecasts/"

	// accuweather API versions

	VERSION_1 Params = "v1/"
)
