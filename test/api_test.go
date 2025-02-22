/*
This file contains tests for NiftyGoGo API functions.
It covers GetCryptoData and error handling via Check.
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
