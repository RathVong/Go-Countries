package country

import (
	"io/ioutil"
	"os"
	"strings"
	"log"
)

const (
	countryFileName = "countries.dat"
	newlineSeparator = "\n"
	commaSeparator = ","
)

func Init() {
	f, err := os.Open(countryFileName)
	defer f.Close()

	if err != nil {
		log.Panicf("Country.init() Error! -> Unable to open countries.dat : %v", err)
	}

	parseDataFile(f.Name())

}

func addCountry(n string, iso string, code int, codeStr string) (country Country) {
	country.Name = n
	country.CountryISO = iso
	country.CountryCode = code
	country.CountryCodeSTR = codeStr
	return
}

func parseDataFile(filename string) (err error){
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panicf("Country.parseDataFile() Error! -> %v", err)
	}

	populateRows(b)

	return
}

func populateRows(b []byte) {
	data := strings.Split(string(b), newlineSeparator)

	for _, row := range data {
		populateFields(row)
	}


}

func populateFields(row string){
	field := strings.Split(row, commaSeparator)
	countries[strings.ToLower(field[0])] = addCountry(field[0],
									 field[1],
							         ConvertStringToInt(field[2]),
									 field[2])
}

