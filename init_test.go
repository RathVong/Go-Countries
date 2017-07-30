package country

import (
	"testing"
	"log"
)

func TestInit(test *testing.T){
	Init()
	log.Println("TestInit() Pass.")
}