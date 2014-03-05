package accuweather

import (
	"accuweather/params"
	"fmt"
	"testing"
)

func Test_SetKeys(t *testing.T) {
	accu, err := SetKeys("fdsaf", "locationKey")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accu.apiKey, accu.locationKey)
	accu, err = SetKeys("apiKey", "")
	if err != nil {
		fmt.Println(err)
	}
	accu, err = SetKeys("", "locationKey")
	if err != nil {
		fmt.Println(err)
	}
	accu, err = SetKeys("", "")
	if err != nil {
		fmt.Println(err)
	}
}
func Test_GetCurrentConditionsUrl(t *testing.T) {
	accu, err := SetKeys("fdsaf", "locationKey")
	if err != nil {
		fmt.Println(err)
	}
	url, _ := accu.GetCurrentConditionsUrl(params.ENVRIOMENT_DEVELOPMENT, params.VERSION_1, params.FORMAT_JSON, params.LANGUAGE_CHINESE, params.DETAILS_TRUE)
	fmt.Println(url)
}
