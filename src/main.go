package main

import (
	"fmt"
	"os"
	"strings"

	"NiftyGoGo/utils"
)

/**
 * NiftyGoGo - A TUI for fetching stock and crypto prices.
 *
 * This app fetches stock prices from Yahoo Finance and crypto prices from Binance.  
 * Users enter a symbol (e.g., "AAPL" for stocks or "BTCUSDT" for crypto),  
 * and the price is fetched and displayed in a terminal UI.
 *
 * Usage:
 *    go run main.go stock   -> Fetch stock prices  
 *    go run main.go crypto  -> Fetch crypto prices  
 */

// URLs holds API endpoints
var URLs = struct {
	Stock  string
	Crypto string
}{
	Stock:  "https://query1.finance.yahoo.com/v7/finance/quote",
	Crypto: "https://api.binance.com/api/v3/ticker/price",
}

func main() {
	// check: ensure mode argument exists
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [stock|crypto]")
		os.Exit(1)
	}

	// convert: normalize input
	mode := strings.ToLower(os.Args[1])

	switch mode {
	case "stock":
		// call: stock TUI using Yahoo Finance
		utils.DispStockPrice(URLs.Stock)

	case "crypto":
		// call: crypto TUI using Binance
		utils.DispCryptoPrice(URLs.Crypto)

	default:
		// error: invalid argument
		fmt.Println("Invalid argument. Use 'stock' or 'crypto'.")
		os.Exit(1)
	}
}

