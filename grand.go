package gogrand

import (
	"net/http"
	"time"
	"unicode"

	"strings"

	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// Exchange is the struct with all the data in it.
type Exchange map[int]struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Members         bool   `json:"members"`
	Sp              int    `json:"sp"`
	BuyAverage      int    `json:"buy_average"`
	BuyQuantity     int    `json:"buy_quantity"`
	SellAverage     int    `json:"sell_average"`
	SellQuantity    int    `json:"sell_quantity"`
	OverallAverage  int    `json:"overall_average"`
	OverallQuantity int    `json:"overall_quantity"`
}

// Item is returned when only a single item is requested.
type Item struct {
	ID              int
	Name            string
	Members         bool
	Sp              int
	BuyAverage      int
	BuyQuantity     int
	SellAverage     int
	SellQuantity    int
	OverallAverage  int
	OverallQuantity int
}

var (
	// Base is the base url.
	Base = "https://storage.googleapis.com/osbuddy-exchange/summary.json"
)

func init() {
	aCache()
}

// GetExchange will return the struct.
func getExchange() (e Exchange, err error) {
	apiClient := http.Client{
		Timeout: 2 * time.Second,
	}

	r, err := apiClient.Get(Base)
	if err != nil {
		return nil, err
	}

	err = jsoniter.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// GetByID will return the item stats based on the ID given.
func GetByID(id int) (i Item, err error) {
	if cached[id].ID != id {
		return Item{}, errors.New("gogrand: item not found")
	}

	i = Item{
		ID:              cached[id].ID,
		Name:            cached[id].Name,
		Members:         cached[id].Members,
		Sp:              cached[id].Sp,
		BuyAverage:      cached[id].BuyAverage,
		BuyQuantity:     cached[id].BuyQuantity,
		SellAverage:     cached[id].SellAverage,
		SellQuantity:    cached[id].SellQuantity,
		OverallAverage:  cached[id].OverallAverage,
		OverallQuantity: cached[id].OverallQuantity,
	}

	return i, nil
}

// GetByName will return the item struct if it can find the item.
func GetByName(name string) (i Item, err error) {

	uc := ucFirst(strings.ToLower(name))

	for key := range cached {
		if cached[key].Name == uc {
			i = Item{
				ID:              cached[key].ID,
				Name:            cached[key].Name,
				Members:         cached[key].Members,
				Sp:              cached[key].Sp,
				BuyAverage:      cached[key].BuyAverage,
				BuyQuantity:     cached[key].BuyQuantity,
				SellAverage:     cached[key].SellAverage,
				SellQuantity:    cached[key].SellQuantity,
				OverallAverage:  cached[key].OverallAverage,
				OverallQuantity: cached[key].OverallQuantity,
			}
			return i, nil
		}

	}
	return Item{}, errors.New("gogrand: item not found")
}

// GetExchange returns the cached json.
func GetExchange() (e Exchange) {
	return cached
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
