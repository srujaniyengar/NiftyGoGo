package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/**
 * GetCryptoData fetches cryptocurrency prices from the given API URL
 * and prints only the requested symbols. It sends a GET request,
 * parses the JSON response, and filters the results.
 */
type PriceData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func GetCryptoData(apiURL string, symbols []string) {
	response, err := http.Get(apiURL)
	Check(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	Check(err)

	var prices []PriceData
	err = json.Unmarshal(body, &prices)
	Check(err)

	for _, p := range prices {
		for _, symbol := range symbols {
			if p.Symbol == symbol {
				fmt.Printf("%s: $%s\n", p.Symbol, p.Price)
			}
		}
	}
}
