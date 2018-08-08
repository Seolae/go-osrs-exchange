package gogrand

import (
	"log"
	"time"
)

var cached Exchange
var UpdatedAt time.Time

func aCache() {
	e, err := getExchange()
	if err != nil {
		log.Panicln("gogrand: Error updating cache for the first time.")
	}

	cached = e
	UpdatedAt = time.Now()
	go func() {
		time.Sleep(time.Second * 150)
		for {
			e, err := getExchange()
			if err != nil {
				log.Println("Error updating cache.")
			}
			cached = e
			UpdatedAt = time.Now()
			time.Sleep(time.Second * 150)
		}
	}()

}

// UpdateCache will manually update the cache.
func UpdateCache() {
	e, err := getExchange()
	if err != nil {
		log.Println("gogrand: Error updating cache.")
	}
	cached = e
}
