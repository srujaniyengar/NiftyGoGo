package main

import "NiftyGoGo/utils"

/**
 * NiftyGoGo - A TUI for fetching stock and crypto prices.
 *
 * This app fetches crypto prices from Binance.
 * The stock feature is coming soon.
 * Users select a mode interactively.
 */

var URLs = struct {
	Crypto string
}{
	Crypto: "https://api.binance.com/api/v3/ticker/price",
}

func main() {
	// call: launch interactive mode selection menu
	utils.DispInteractiveChoice(URLs.Crypto)
}
