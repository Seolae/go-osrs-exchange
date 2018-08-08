package gogrand

import (
	"log"
	"testing"
)

func TestExchange(t *testing.T) {
	e := GetExchange()

	log.Println(e[2])
}

func TestGetByID(t *testing.T) {
	i, err := GetByID(1)
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(i)
}

func TestGetByName(t *testing.T) {
	i, err := GetByName("yew shortbow (u)")
	if err != nil {
		t.Error(err.Error())
	}

	log.Println(i)
}
