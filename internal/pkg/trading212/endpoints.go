package trading212

type APIEndpoint string

const version APIEndpoint = "api/v0"

const endpointBase = version + "/equity"

//nolint:godoclint,godot
//goland:noinspection GoCommentStart
const (
	// Account
	GetAccountSummary = endpointBase + "/account/summary"
	// Instruments
	GetExchangesMetadata       = endpointBase + "/metadata/exchanges"
	GetAllAvailableInstruments = endpointBase + "/metadata/instruments"
	// Orders
	GetAllPendingOrders = endpointBase + "/orders"
	GetPendingOrderByID = endpointBase + "/orders" // + /{id}
	PlaceLimitOrder     = endpointBase + "/orders/limit"
	PlaceMarketOrder    = endpointBase + "/orders/market"
	PlaceStopOrder      = endpointBase + "/orders/stop"
	PlaceStopLimitOrder = endpointBase + "/orders/stop_limit"
	CancelOrder         = endpointBase + "/orders" // + /{id}
	// Positions
	GetAllPositions = endpointBase + "/positions"
	// Historical events
	GetDividends        = endpointBase + "/history/dividends"
	GetHistoricalOrders = endpointBase + "/history/orders"
	GetTransactions     = endpointBase + "/history/transactions"
	ListReports         = endpointBase + "/history/exports"
	RequestReport       = endpointBase + "/history/exports"
	// Pies
	GetAllPies   = endpointBase + "/pies"
	CreatePie    = endpointBase + "/pies"
	DeletePie    = endpointBase + "/pies" // + /{id}
	FetchPie     = endpointBase + "/pies" // + /{id}
	UpdatePie    = endpointBase + "/pies" // + /{id}
	DuplicatePie = endpointBase + "/pies" // + /{id}/duplicate
)
