# Go-Countries
A Simple Countries Utilities For GO



This utility is meant for quick country selection activites. Enjoy!





Installation
-------------------------------
    go get "github.com/RathVong/Go-Countries"
 
    import "github.com/RathVong/Go-Countries"


Setup
-------------------------------

First initialize the countries.


    func main(){
      country.init()
    }
    
    
Get a country by name ->

    country := country.get("Canada")
  
Get a country by ISO ->

    country := country.FindByISO("CA")

Get country by code ->

    country := country.FindByCode(1)

To get the entire list of country names ->

    names := country.GetCountryNames()
    

Data structure
--------------------------------

  Incase you are planning on marshalling to json, here are the
  struct and keys.

           type Country struct {
	            Name           string `json:"name"`
	            CountryISO     string `json:"country_iso"`
	            CountryCode    int    `json:"country_code"`
	            CountryCodeSTR string `json:"country_code_str"`
           }


  
  

