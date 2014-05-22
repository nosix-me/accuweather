package accuweather

import (
	"log"
	"testing"
)

func Test_SetKeys(t *testing.T) {
	accu, err := SetKeys("fdsaf", "locationKey")
	if err != nil {
		log.Println(err)
	}
	log.Println(accu.apiKey, accu.locationKey)
	accu, err = SetKeys("apiKey", "")
	if err != nil {
		log.Println(err)
	}
	accu, err = SetKeys("", "locationKey")
	if err != nil {
		log.Println(err)
	}
	accu, err = SetKeys("", "")
	if err != nil {
		log.Println(err)
	}
}
func Test_GetCurrentConditionsUrl(t *testing.T) {
	accu, err := SetKeys("fdsaf", "locationKey")
	if err != nil {
		log.Println(err)
	}
	url, _ := accu.GetCurrentConditionsUrl(ENVRIOMENT_DEVELOPMENT, VERSION_1, FORMAT_JSON, LANGUAGE_CHINESE, DETAILS_TRUE)
	log.Println(url)
}
