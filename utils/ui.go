package utils

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
)

func white() string { // add: white color
	return "\033[37m"
}

/**
 * TUI for displaying stock and cryptocurrency prices using Bubble Tea.
 * Users can enter a crypto symbol, and the corresponding price is fetched and displayed.
 *
 * - DispStockPrice(stockUrl string): Displays stock prices.
 * - DispCryptoPrice(cryptoUrl string): Displays cryptocurrency prices.
 *
 * The UI updates dynamically based on user input and fetches data accordingly.
 */

type menuModel struct {
	selected int
}

func (m menuModel) Init() tea.Cmd { // init:
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { // update:
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" { // fix: exit
			return m, tea.Quit
		}
		switch msg.String() {
		case "left", "h": // calls: move left
			if m.selected > 0 {
				m.selected--
			}
		case "right", "l": // calls: move right
			if m.selected < 1 {
				m.selected++
			}
		case "enter": // calls: confirm
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m menuModel) View() string { // view:
	clear := "\033[H\033[2J"
	var s string
	if m.selected == 0 {
		s = fmt.Sprintf(" > [%sStock%s]    [%sCrypto%s]\n", cyan, reset, white(), reset)
	} else {
		s = fmt.Sprintf("   [%sStock%s]  > [%sCrypto%s]\n", white(), reset, cyan, reset)
	}
	return clear + "Select Mode:\n\n" + s + "\nUse arrow keys (or h/l) and press Enter. (Press ctrl+c to exit)"
}

type inputModel struct {
	symbol string
	price  string
	err    error
	url    string
	mode   string
}

func (m inputModel) Init() tea.Cmd { // init:
	return nil
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { // update:
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" { // fix: exit
			return m, tea.Quit
		}
		switch msg.String() {
		case "enter": // calls: fetch price
			if m.mode == "crypto" {
				priceMap, err := GetCryptoData(m.url, []string{strings.ToUpper(m.symbol)})
				if err != nil {
					m.err = err
				} else {
					if p, ok := priceMap[strings.ToUpper(m.symbol)]; ok {
						m.price = "$" + p
					} else {
						m.err = fmt.Errorf("symbol not found: %s", m.symbol)
					}
				}
			}
			return m, nil
		case "backspace": // fix: remove char
			if len(m.symbol) > 0 {
				m.symbol = m.symbol[:len(m.symbol)-1]
			}
		default: // calls: append char
			m.symbol += msg.String()
		}
	}
	return m, nil
}

func (m inputModel) View() string { // view:
	clear := "\033[H\033[2J"
	if m.mode == "crypto" {
		popularCoins := []string{
			"BTCUSDT",
			"ETHUSDT",
			"BNBUSDT",
			"ADAUSDT",
			"XRPUSDT",
			"DOGEUSDT",
			"SOLUSDT",
			"DOTUSDT",
			"LTCUSDT",
			"LINKUSDT",
		}
		coinsList := ""
		for _, coin := range popularCoins {
			coinsList += coin + "\n"
		}
		// add: prompt message updated with instruction
		output := fmt.Sprintf("%s%sPopular coins:%s\n%s\nEnter crypto symbol (Press enter to see latest price): %s\n(Press ctrl+c to exit)", clear, bold, reset, coinsList, cyan+m.symbol+reset)
		if m.price != "" {
			output += fmt.Sprintf("\n%sCrypto Price for %s: %s%s", green, strings.ToUpper(m.symbol), m.price, reset)
		}
		if m.err != nil {
			output += fmt.Sprintf("\n%sError: %v%s", red, m.err, reset)
		}
		return output
	}
	return clear + "Unknown mode\n(Press ctrl+c to exit)"
}

type comingSoonModel struct{}

func (m comingSoonModel) Init() tea.Cmd { // init:
	return nil
}

func (m comingSoonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { // update:
	if k, ok := msg.(tea.KeyMsg); ok {
		if k.String() == "ctrl+c" { // fix: exit
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m comingSoonModel) View() string { // view:
	clear := "\033[H\033[2J"
	return clear + yellow + "Stock feature coming soon: stay tuned\n(Press ctrl+c to exit)" + reset
}

func DispInteractiveChoice(cryptoUrl string) { // calls:
	menu := menuModel{selected: 0}
	p := tea.NewProgram(menu)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	selection := finalModel.(menuModel)
	if selection.selected == 0 {
		p2 := tea.NewProgram(comingSoonModel{})
		_, err := p2.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
		cryptoModel := inputModel{
			url:  cryptoUrl,
			mode: "crypto",
		}
		p3 := tea.NewProgram(cryptoModel)
		_, err := p3.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
