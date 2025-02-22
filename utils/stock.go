
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

func GetStockPrice(url, symbol string) (float64, error) {
	// fetch: from Yahoo Finance API
	resp, err := http.Get(fmt.Sprintf("%s?symbols=%s", url, symbol))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// decode: JSON response into QuoteResponse struct
	var data QuoteResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}

	// check: if the response contains stock data
	if len(data.QuoteResponse.Result) == 0 {
		return 0, fmt.Errorf("no data found for symbol: %s", symbol)
	}

	// return: stock price of the requested symbol
	return data.QuoteResponse.Result[0].RegularMarketPrice, nil
}


