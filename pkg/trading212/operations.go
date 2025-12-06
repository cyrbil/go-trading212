// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"io"

	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

// helper function for the operations.
func runOperation[T any](api requestMaker, method string, endpoint internal.APIEndpoint, body any) *Response[T] {
	var requestBodyReader io.Reader
	if body != nil {
		requestBodyReader = newJSONBody(body)
	}

	request, err := api.NewRequest(method, endpoint, requestBodyReader)
	if err != nil {
		return &Response[T]{request: request, raw: nil, err: err}
	}

	data, err := request.Do()
	if err != nil {
		return &Response[T]{request: request, raw: data, err: err}
	}

	return &Response[T]{request: request, raw: data, err: nil}
}

// operations regroups all available operations.
type operations struct {
	// Account operations.
	// Access fundamental information about your trading account.
	// Retrieve details such as your account ID, currency, and current cash balance.
	// See: https://docs.trading212.com/api/accounts
	Account accountOperations

	// Instruments operations.
	// Discover what you can trade. These endpoints provide comprehensive lists of all tradable instruments
	// and the exchanges they belong to, including details like tickers and trading hours.
	// See: https://docs.trading212.com/api/instruments
	Instruments instrumentsOperations

	// Orders operations
	// Place, monitor, and cancel equity trade orders.
	// This section provides the core functionality for
	// programmatically executing your trading strategies for stocks and ETFs.
	// See: https://docs.trading212.com/api/orders
	Orders ordersOperations

	// Positions operations
	// Get a real-time overview of all your open positions, including quantity, average price, and current profit or loss.
	// See: https://docs.trading212.com/api/positions
	Positions positionsOperations

	// HistoricalEvents operations
	// Review your account's trading history. Access detailed records of past
	// orders, dividend payments, and cash transactions, or generate downloadable
	// CSV reports for analysis and record-keeping.
	// See: https://docs.trading212.com/api/historical-events
	HistoricalEvents historicalEventsOperations

	// Pies operations
	// Manage your investment Pies. Use these endpoints to create, view, update,
	// and delete your custom portfolios, making automated and diversified investing simple.
	//
	// Deprecation notice: The current state of the Pies API,
	// while still operational, won't be further supported and updated.
	// See: https://docs.trading212.com/api/pies-(deprecated)
	Pies piesOperations
}
