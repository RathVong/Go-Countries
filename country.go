package country

import "strings"

type Country struct {
	Name string
	CountryISO string
	CountryCode int
	CountryCodeSTR string
}


var countries map[string] Country = make(map[string]Country, 233)



func  Get(name string) (country Country){
	if val, ok := countries[strings.ToLower(name)]; ok {
		return val
	}

	return
}

func  GetCountryNames() (names []string){
	for _, v := range countries{
		names = append(names, v.Name)
	}

	return
}

// When querying for Canada country code 1 it will return US only.
func  FindByCode(code int)(country Country){
	for _, v := range countries {
		if v.CountryCode == code && v.CountryISO == "us"  {
			return v
		}

	}

	return
}

func  FindByISO(iso string) (country Country){
	for _, v := range countries {
		if strings.ToLower(v.CountryISO) == strings.ToLower(iso) {
			return v
		}
	}

	return
}


