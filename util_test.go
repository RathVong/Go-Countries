package country

import (
	"testing"
	"log"
)

func TestConvertStringToInt(t *testing.T) {
	numberString := "10"

	n := ConvertStringToInt(numberString)

	if n != 10 {
		t.Errorf("Converting 10 string to number returns %v", n)
	}

	log.Println("TestConvertStringToInt() Pass.")

}