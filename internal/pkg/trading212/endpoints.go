// Package trading212 github.com/cyrbil/go-trading212
package trading212

// APIEndpoint type
type APIEndpoint string

const version APIEndpoint = "/api/v0"

const endpointBase = version + "/equity"

//nolint:godoclint,godot
//goland:noinspection GoCommentStart
const (
	// Account

	// GetAccountSummary endpoint
	GetAccountSummary = endpointBase + "/account/summary"

	// Instruments

	// GetExchangesMetadata endpoint
	GetExchangesMetadata = endpointBase + "/metadata/exchanges"
	// GetAllAvailableInstruments endpoint
	GetAllAvailableInstruments = endpointBase + "/metadata/instruments"

	// Orders

	// GetAllPendingOrders endpoint
	GetAllPendingOrders = endpointBase + "/orders"
	// GetPendingOrderByID endpoint
	GetPendingOrderByID = endpointBase + "/orders" // + /{id}
	// PlaceLimitOrder endpoint
	PlaceLimitOrder = endpointBase + "/orders/limit"
	// PlaceMarketOrder endpoint
	PlaceMarketOrder = endpointBase + "/orders/market"
	// PlaceStopOrder endpoint
	PlaceStopOrder = endpointBase + "/orders/stop"
	// PlaceStopLimitOrder endpoint
	PlaceStopLimitOrder = endpointBase + "/orders/stop_limit"
	// CancelOrder endpoint
	CancelOrder = endpointBase + "/orders" // + /{id}

	// Positions

	// GetAllPositions endpoint
	GetAllPositions = endpointBase + "/positions"

	// Historical events

	// GetDividends endpoint
	GetDividends = endpointBase + "/history/dividends"
	// GetHistoricalOrders endpoint
	GetHistoricalOrders = endpointBase + "/history/orders"
	// GetTransactions endpoint
	GetTransactions = endpointBase + "/history/transactions"
	// ListReports endpoint
	ListReports = endpointBase + "/history/exports"
	// RequestReport endpoint
	RequestReport = endpointBase + "/history/exports"

	// Pies

	// GetAllPies endpoint
	GetAllPies = endpointBase + "/pies"
	// CreatePie endpoint
	CreatePie = endpointBase + "/pies"
	// DeletePie endpoint
	DeletePie = endpointBase + "/pies" // + /{id}
	// FetchPie endpoint
	FetchPie = endpointBase + "/pies" // + /{id}
	// UpdatePie endpoint
	UpdatePie = endpointBase + "/pies" // + /{id}
	// DuplicatePie endpoint
	DuplicatePie = endpointBase + "/pies" // + /{id}/duplicate
)
