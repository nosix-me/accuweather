// accuweather project accuweather.go
//auther:Yaxiang He
package accuweather

import (
	"accuweather/params"
	"errors"
)

type Accuweather struct {
	apiKey      string
	locationKey string
}

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
func (accu *Accuweather) GetCurrentConditionsUrl(enviroment params.Params, version params.Params, format params.Params, language params.Params, details params.Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys params error!")
	}
	return string(enviroment) + string(params.TYPES_CURRENTCONDITIONS) + string(version) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get current time prev 6 hours history data
func (accu *Accuweather) GetHistoricalCurrentConditionsUrl(enviroment params.Params, version params.Params, format params.Params, locationKey string, apiKey string, language params.Params, details params.Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys params error!")
	}
	return string(enviroment) + string(params.TYPES_CURRENTCONDITIONS) + string(version) + accu.locationKey + "/historical" + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get future weather data for hourly
func (accu *Accuweather) GetForecastHourly(enviroment params.Params, forecasthourly params.Params, version params.Params, format params.Params, locationKey string, apiKey string, language params.Params, details params.Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys params error!")
	}
	return string(enviroment) + string(params.TYPES_FORECASTS) + string(version) + string(params.FROECAST_HOURLY) + string(forecasthourly) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}

//Get future weather data for daily
func (accu *Accuweather) GetForecastDaily(enviroment params.Params, forecastdaily params.Params, version params.Params, format params.Params, locationKey string, apiKey string, language params.Params, details params.Params) (string, error) {
	if len(accu.locationKey) == 0 || len(accu.apiKey) == 0 {
		return "", errors.New("Accuweather keys params error!")
	}
	return string(enviroment) + string(params.TYPES_FORECASTS) + string(version) + string(params.FROECAST_DAILY) + string(forecastdaily) + accu.locationKey + string(format) + "apiKey=" + accu.apiKey + "&language=" + string(language) + "&details=" + string(details), nil
}
