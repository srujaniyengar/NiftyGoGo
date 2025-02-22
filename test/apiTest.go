/*
This file contains tests for NiftyGoGo API functions.
It covers GetCryptoData, GetStockPrice, and error handling via Check.
Tests simulate API endpoints with httptest.
*/

package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"NiftyGoGo/utils"
)

// TestGetCryptoData tests the GetCryptoData function.
func TestGetCryptoData(t *testing.T) {
	// create: test server for crypto API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetch: simulated crypto JSON data
		data := `[{"symbol":"BTCUSDT","price":"30000"},{"symbol":"ETHUSDT","price":"2000"}]`
		w.Header().Set("Content-Type", "application/json")
		// call: write response
		w.Write([]byte(data))
	}))
	defer ts.Close()

	// call: test for existing symbol
	symbols := []string{"BTCUSDT"}
	priceMap, err := utils.GetCryptoData(ts.URL, symbols)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// check: BTCUSDT price is correct
	if price, ok := priceMap["BTCUSDT"]; !ok || price != "30000" {
		t.Errorf("Expected BTCUSDT price '30000', got %v", price)
	}

	// call: test for non-existent symbol
	symbols = []string{"NONEXISTENT"}
	_, err = utils.GetCryptoData(ts.URL, symbols)
	if err == nil {
		t.Error("Expected error for non-existent symbol, got nil")
	}
}

// TestGetStockPrice tests the GetStockPrice function.
func TestGetStockPrice(t *testing.T) {
	// create: test server for stock API
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fetch: retrieve query symbol
		symbol := r.URL.Query().Get("symbols")
		if symbol == "AAPL" {
			// call: valid response for AAPL
			jsonResponse := `{"quoteResponse": {"result": [{"symbol": "AAPL", "regularMarketPrice": 150.50}]}}`
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(jsonResponse))
		} else {
			// call: empty response for other symbols
			jsonResponse := `{"quoteResponse": {"result": []}}`
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(jsonResponse))
		}
	}))
	defer ts.Close()

	// call: test valid symbol AAPL
	price, err := utils.GetStockPrice(ts.URL, "AAPL")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// check: price equals 150.50
	if price != 150.50 {
		t.Errorf("Expected price 150.50, got %v", price)
	}

	// call: test invalid symbol
	_, err = utils.GetStockPrice(ts.URL, "INVALID")
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}

// TestCheck tests the Check function.
func TestCheck(t *testing.T) {
	// call: Check should panic for non-nil error
	defer func() {
		// check: panic occurred
		if r := recover(); r == nil {
			t.Error("Expected panic for non-nil error, but function did not panic")
		}
	}()
	utils.Check(fmt.Errorf("test error"))
}
