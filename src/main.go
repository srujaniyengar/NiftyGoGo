package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"NiftyGoGo/utils"
)

/*
This main file prompts the user to choose a mode.
It then launches the appropriate TUI for stock or crypto prices.
*/

var URLs = struct {
	Stock  string
	Crypto string
}{
	Stock:  "https://query1.finance.yahoo.com/v7/finance/quote",
	Crypto: "https://api.binance.com/api/v3/ticker/price",
}

func main() {
	// prompt: ask the user for input
	fmt.Println("Enter 'stock' for Stock Prices or 'crypto' for Crypto Prices:")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	// normalize: remove extra spaces and lower-case the input
	mode := strings.TrimSpace(strings.ToLower(choice))

	switch mode {
	case "stock":
		// call: start stock TUI
		utils.DispStockPrice(URLs.Stock)
	case "crypto":
		// call: start crypto TUI
		utils.DispCryptoPrice(URLs.Crypto)
	default:
		// error: invalid input
		fmt.Println("Invalid choice. Please run the program again and choose either 'stock' or 'crypto'.")
		os.Exit(1)
	}
}
