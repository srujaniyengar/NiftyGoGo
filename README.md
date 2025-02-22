# NiftyGoGo  

NiftyGoGo is a terminal UI (TUI) application for fetching cryptocurrency prices from Binance. The stock price feature is coming soon.  

## Features  

- Interactive mode selection (Stock / Crypto)  
- Displays a list of popular crypto coins vertically  
- Fetches the latest crypto prices from Binance  
- **Stock feature: Coming soon**  
- Built using [Bubble Tea](https://github.com/charmbracelet/bubbletea)  

## Requirements  

- Go 1.16 or later  
## Installation & Usage  

1. Clone the repository.  
2. In the project root, run:  

   ```sh
   ./run.sh
   ```
 3. Alternative command:
 ```sh
 sh run.sh
```
# How It Works

## Mode Selection  
Use arrow keys (or `h/l`) to choose between **Stock** and **Crypto**.

- **Stock Mode:**  
  Displays `"Stock feature coming soon: stay tuned"`.  

- **Crypto Mode:**  
  - Displays a list of popular coins.  
  - Enter a crypto symbol (e.g., `BTCUSDT`, `ETHUSDT`) to see the latest price.  
  - Press `Enter` to refresh and get the latest price.  

## Exit  
Press `Ctrl+C` at any time to exit.

# Data Source  

This application retrieves cryptocurrency price data from the **Binance API**.  
All rights to the data belong to **Binance**.  
For more details, visit [Binance API Documentation](https://binance-docs.github.io/apidocs/spot/en/).

