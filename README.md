# go-trading212

[![Go Reference][go-reference-badge]][go-reference]
[![License][license-badge]][license]
[![Go Version][go-version-badge]][go-version]
[![coded-by-badge][coded-by-badge]][repo-commits]

[![Release Status][release-badge]][release]
[![Coverage Status][coverage-badge]][coverage]
[![Testing Status][testing-badge]][testing]

[coverage]: https://github.com/cyrbil/go-trading212/actions/workflows/coverage.yml
[coverage-badge]: https://raw.githubusercontent.com/cyrbil/go-trading212/badges/.badges/main/coverage.svg?branch=main&event=push
[testing]: https://github.com/cyrbil/go-trading212/actions/workflows/testing.yml
[testing-badge]: https://github.com/cyrbil/go-trading212/actions/workflows/testing.yml/badge.svg?branch=main&event=schedule
[release]: https://github.com/cyrbil/go-trading212/actions/workflows/release.yml
[release-badge]: https://github.com/cyrbil/go-trading212/actions/workflows/release.yml/badge.svg?branch=main&event=push


A comprehensive Go client library for interacting with the [Trading212 Rest API][trading212-docs].
This library provides a type-safe, idiomatic Go interface for managing your Trading212 account, 
placing orders, monitoring positions, and accessing historical trading data.


## Features

- 🔐 **Secure Authentication** - Built-in support for API key and secret authentication with secure string handling
- 📊 **Account Management** - Retrieve account summaries, cash balances, and investment metrics
- 📈 **Order Management** - Place and manage market, limit, stop, and stop-limit orders
- 🔍 **Instrument Discovery** - Browse available instruments and exchange metadata
- 📍 **Position Tracking** - Monitor open positions with real-time profit/loss data
- 📜 **Historical Data** - Access trading history, dividends, transactions, and generate CSV reports
- 🥧 **Pies Management** - Manage investment pies (deprecated API)
- ⚡ **Rate Limiting** - Built-in rate limit handling to respect API constraints
- 🎯 **Type Safety** - Full type safety with Go's strong typing system
- 🔄 **Iterator Support** - Modern iterator-based API for streaming large datasets


## Installation

```bash
go get github.com/cyrbil/go-trading212
```


## Quick Start


### Basic Setup

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/cyrbil/go-trading212/pkg/trading212"
)

func main() {
    // Initialize the API client
    api := trading212.NewAPILive(
        "your-api-key",
        "your-api-secret",
    )
    
    // Get account summary
    summary, err := api.Account.GetAccountSummary()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Account ID: %d\n", summary.Id)
    fmt.Printf("Currency: %s\n", summary.Currency)
    fmt.Printf("Total Value: %.2f\n", summary.TotalValue)
    fmt.Printf("Available to Trade: %.2f\n", summary.Cash.AvailableToTrade)
}
```


### Get All Open Positions

```go
// Retrieve all open positions
positions, err := api.Positions.GetAllPositions()
if err != nil {
    log.Fatal(err)
}

for position := range positions {
    fmt.Printf("Position: %s - Quantity: %.2f - P/L: %.2f\n",
        position.Ticker,
        position.Quantity,
        position.UnrealizedProfitLoss,
    )
}
```


## API Overview

The library is organized into logical operation groups:

### Account Operations
- `GetAccountSummary()` - Get account details, cash balance, and investment metrics

### Instrument Operations
- `GetExchangesMetadata()` - Get all exchanges and their working schedules
- `GetAllAvailableInstruments()` - Get all tradable instruments

### Order Operations
- `PlaceMarketOrder()` - Place a market order
- `PlaceLimitOrder()` - Place a limit order
- `PlaceStopOrder()` - Place a stop order
- `PlaceStopLimitOrder()` - Place a stop-limit order
- `GetAllPendingOrders()` - Get all active orders
- `GetPendingOrderByID()` - Get a specific pending order
- `CancelOrder()` - Cancel an active order

### Position Operations
- `GetAllPositions()` - Get all open positions

### Historical Events Operations
- `GetPaidOutDividends()` - Get dividend payment history
- `GetHistoricalOrders()` - Get historical order fills
- `GetTransactions()` - Get account transactions
- `ListReports()` - List available CSV reports
- `RequestReport()` - Request a new CSV report

### Pies Operations (Deprecated)
- `FetchAllPies()` - Get all investment pies
- `CreatePie()` - Create a new pie
- `FetchPie()` - Get pie details
- `UpdatePie()` - Update a pie
- `DeletePie()` - Delete a pie
- `DuplicatePies()` - Duplicate a pie

## Configuration


### API Domains

The library also supports demo or any trading212 environments:

```go
// Demo environment (for testing)
api := trading212.NewAPI(
    trading212.APIDomainDemo,
    apiKey,
    apiSecret,
)

// Custom environment
api := trading212.NewAPI(
    trading212.APIDomain("api.domain"),
    apiKey,
    apiSecret,
)
```


### Secure String

The library uses a `SecureString` type for API secrets to prevent accidental logging of sensitive credentials:

```go
apiSecret := trading212.SecureString("your-secret-key")
// When printed, this will show "[REDACTED]" instead of the actual value
fmt.Println(apiSecret) // Output: [REDACTED]
```


## Error Handling

All operations return errors that should be checked:

```go
summary, err := api.Account().GetAccountSummary()
if err != nil {
    // Handle error appropriately
    log.Printf("Failed to get account summary: %v", err)
    return
}
// Use summary...
```


## Rate Limiting

The library includes built-in rate limiting support.
Rate limits are automatically tracked per endpoint to ensure compliance with Trading212 API constraints.


## Requirements

- Go 1.23 or higher
- A Trading212 account with API access enabled
- Valid API key and secret


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request


## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.


## Disclaimer

This library is not affiliated with, endorsed by, or sponsored by Trading212. Use at your own risk. 


[go-reference-badge]: https://pkg.go.dev/badge/github.com/cyrbil/go-trading212.svg
[go-reference]: https://pkg.go.dev/github.com/cyrbil/go-trading212
[license-badge]: https://img.shields.io/badge/license-GPLv3-blue.svg
[license]: ./LICENSE
[go-version-badge]: https://img.shields.io/badge/go-1.23+-00ADD8.svg
[go-version]: https://golang.org
[trading212-docs]: https://docs.trading212.com/api
[coded-by-badge]: https://img.shields.io/badge/coded%20by-humans%20%F0%9F%92%96-blue?style=social
[repo-commits]: https://github.com/cyrbil/go-trading212/commits/main/
