
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * stock.go - Fetch stock prices using Yahoo Finance API.
 * Steps: fetch → decode → validate → return
 */

// QuoteResponse defines the response structure from Yahoo Finance API
type QuoteResponse struct {
	QuoteResponse struct {
		Result []struct {
			Symbol             string  `json:"symbol"`
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

// GetStockPrice fetches the stock price for a given symbol from the provided API URL
func GetStockPrice(apiURL, symbol string) float64 {
	// fetch: GET request to Yahoo Finance API
	url := fmt.Sprintf("%s?symbols=%s", apiURL, symbol)
	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	// decode: JSON response into QuoteResponse struct
	var data QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	Check(err)

	// validate: check if response contains stock data
	if len(data.QuoteResponse.Result) == 0 {
		panic(fmt.Sprintf("no data found for symbol: %s", symbol))
	}

	// return: stock price of the requested symbol
	return data.QuoteResponse.Result[0].RegularMarketPrice
}

