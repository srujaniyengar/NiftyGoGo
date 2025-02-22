
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// QuoteResponse defines the response structure from Yahoo Finance API
type QuoteResponse struct {
	QuoteResponse struct {
		Result []struct {
			Symbol             string  `json:"symbol"`
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

// GetStockPrice fetches the stock price for a given symbol
func GetStockPrice(symbol string) float64 {
	// fetch: from Yahoo Finance API
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s", symbol)
	resp, err := http.Get(url)
	Check(err) // check: handle request error
	defer resp.Body.Close()

	// decode: JSON response into QuoteResponse struct
	var data QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	Check(err) // check: handle decoding error

	// check: if the response contains stock data
	if len(data.QuoteResponse.Result) == 0 {
		panic(fmt.Sprintf("no data found for symbol: %s", symbol))
	}

	// return: stock price of the requested symbol
	return data.QuoteResponse.Result[0].RegularMarketPrice
}

