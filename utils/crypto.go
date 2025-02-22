package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/**
 * crypto.go - Fetch cryptocurrency prices from an API.
 * Steps: fetch → read → parse → filter
 */

type PriceData struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// GetCryptoData fetches cryptocurrency prices from the given API URL
// and returns a map of symbols to their prices.
func GetCryptoData(apiURL string, symbols []string) (map[string]string, error) {
	// fetch: GET request to API
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read: API response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse: JSON response into PriceData struct
	var prices []PriceData
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, err
	}

	// filter: Store only requested symbols in a map
	priceMap := make(map[string]string)
	for _, p := range prices {
		for _, s := range symbols {
			if p.Symbol == s {
				priceMap[p.Symbol] = p.Price
			}
		}
	}

	// return: Error if no matching symbols found
	if len(priceMap) == 0 {
		return nil, errors.New("no matching symbols found")
	}

	return priceMap, nil
}
