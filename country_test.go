package country

import (
	"testing"
	"log"
)

func TestCountry_Get(t *testing.T) {
	Init()
	country := Get("canada")

	if country.Name != "Canada"{
		t.Errorf("Searching for Canada, found -> %+v", country)
		t.FailNow()
	}

	log.Println("TestCountry_Get() Pass.")
}

func TestCountry_FindByCode(t *testing.T) {
	Init()
	country := FindByCode(1)

	if country.CountryCode != 1 {
		t.Errorf("Searching for country with code 1, found %+v", country )

	} else {

		log.Println("TestCountry_FindByCode() Pass.")
	}
}

func TestCountry_FindByISO(t *testing.T) {
	Init()
	country := FindByISO("CA")

	if country.CountryISO != "ca" {
		t.Errorf("Searching for country with ISO code CA, found %+v", country)
	} else {
		log.Println("TestCountry_FindByISO() Pass.")
	}
}

func TestCountry_GetCountryNames(t *testing.T) {
	Init()
	names := GetCountryNames()

	var n = len(names)

	if n != 233 {
		t.Errorf("TestCountry_GetCountryNames() -> Returned %v", n)
	} else if !validateCountryNames(names) {
		t.Errorf("TestCountry_GetCountryNames() Country names are incorrect")
	}else {

		log.Println("TestCountry_GetCountryNames() Pass.")
	}


}

func validateCountryNames (names []string) (isValidated bool){
	for _, v := range names {

		if len(v) == 0 {
			return false
		}

	}

	return true
}