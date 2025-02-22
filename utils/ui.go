package utils

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

/**
 * TUI for displaying stock and cryptocurrency prices using Bubble Tea.
 * Users can enter a stock/crypto symbol, and the corresponding price is fetched and displayed.
 *
 * - DispStockPrice(stockUrl string): Displays stock prices.
 * - DispCryptoPrice(cryptoUrl string): Displays cryptocurrency prices.
 *
 * The UI updates dynamically based on user input and fetches data accordingly.
 */

// model represents the TUI state
type model struct {
	symbol string
	price  string
	err    error
	url    string
}

// Init initializes the TUI (not used here)
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles user input and updates the model state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// fetch: stock/crypto price
			price, err := GetStockPrice(m.url, strings.ToUpper(m.symbol))
			if err != nil {
				m.err = err
			} else {
				m.price = fmt.Sprintf("$%.2f", price)
			}
			return m, nil

		case "ctrl+c":
			// quit: on ctrl+c
			return m, tea.Quit

		default:
			// read: user input
			m.symbol += msg.String()
		}
	}
	return m, nil
}

// View renders the UI
func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPress any key to exit.", m.err)
	}
	if m.price != "" {
		return fmt.Sprintf("Stock/Crypto Price for %s: %s\nPress any key to exit.", m.symbol, m.price)
	}
	return fmt.Sprintf("Enter stock/crypto symbol: %s", m.symbol)
}

// DispStockPrice starts the TUI for stock prices
func DispStockPrice(stockUrl string) {
	p := tea.NewProgram(model{url: stockUrl})
	_, err := p.Run() // p.Run() returns (tea.Model, error)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// DispCryptoPrice starts the TUI for cryptocurrency prices
func DispCryptoPrice(cryptoUrl string) {
	p := tea.NewProgram(model{url: cryptoUrl})
	_, err := p.Run() // p.Run() returns (tea.Model, error)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

