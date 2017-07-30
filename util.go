package country

import (
	"strconv"
	"log"
)

func ConvertStringToInt(s string) int {

	if len(s) == 0 {
		return 0
	}

	c, err := strconv.ParseInt(s, 10, 64)
	defer func() {
		if err != nil {
			log.Panicf("Country.ConvertStringToInt -> Error: ", err)
		}
	}()



	return int(c)

}