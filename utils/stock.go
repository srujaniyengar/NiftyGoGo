package utils

import (
	"encoding/json"
	"errors"
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
func GetStockPrice(apiURL, symbol string) (float64, error) {
	// fetch: GET request to Yahoo Finance API
	url := fmt.Sprintf("%s?symbols=%s", apiURL, symbol)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err // check: request error
	}
	defer resp.Body.Close()

	// decode: JSON response into QuoteResponse struct
	var data QuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err // check: decoding error
	}

	// validate: check if response contains stock data
	if len(data.QuoteResponse.Result) == 0 {
		return 0, errors.New("no data found for symbol: " + symbol)
	}

	// return: stock price of the requested symbol
	return data.QuoteResponse.Result[0].RegularMarketPrice, nil
}

