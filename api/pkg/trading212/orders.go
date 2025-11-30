package trading212

import "time"

type Order struct {
	// The ISO 8601 formatted date of when the order was created.
	CreatedAt time.Time `json:"createdAt"`
	// The currency used for the order in ISO 4217 format.
	Currency string `json:"currency"`
	// If true, the order is eligible for execution outside regular trading hours.
	ExtendedHours bool `json:"extendedHours"`
	// The number of shares that have been successfully executed. Applicable to quantity orders.
	FilledQuantity float64 `json:"filledQuantity"`
	// The total monetary value of the executed portion of the order. Applicable to orders placed by value.
	// Note: Placing orders by value is not currently supported via the API.
	FilledValue float64 `json:"filledValue"`
	// A unique, system-generated identifier for the order.
	Id uint `json:"id"`
	// How the order was initiated.
	// Enum: "API", "IOS", "ANDROID", "WEB", "SYSTEM", "AUTOINVEST"
	InitiatedFrom string `json:"initiatedFrom"`
	// Instrument information as given by /instruments endpoint.
	Instrument struct {
		// Instrument currency in ISO 4217 format.
		Currency string `json:"currency"`
		// ISIN of the instrument.
		Isin string `json:"isin"`
		// Name of the instrument.
		Name string `json:"name"`
		// Unique instrument identifier. (Example: "AAPL_US_EQ")
		Ticker string `json:"ticker"`
	} `json:"instrument"`
	// Applicable to LIMIT and STOP_LIMIT orders.
	LimitPrice float64 `json:"limitPrice"`
	// The total number of shares requested. Applicable to quantity orders.
	Quantity float64 `json:"quantity"`
	// Indicates whether the order is BUY or SELL.
	Side string `json:"side"`
	// The current state of the order in its lifecycle.
	// Enum: "LOCAL", "UNCONFIRMED", "CONFIRMED", "NEW", "CANCELLING", "CANCELLED", "PARTIALLY_FILLED",
	// "FILLED", "REJECTED", "REPLACING", "REPLACED"
	Status string `json:"status"`
	// Applicable to STOP and STOP_LIMIT orders.
	StopPrice float64 `json:"stopPrice"`
	// The strategy used to place the order, either by QUANTITY or VALUE.
	// The API currently only supports placing orders by QUANTITY.
	//  Enum: "QUANTITY", "VALUE"
	Strategy string `json:"strategy"`
	// Unique instrument identifier. Get from the /instruments endpoint
	Ticker string `json:"ticker"`
	// Specifies how long the order remains active:
	//  - DAY: The order will automatically expire if not executed by midnight
	//    in the time zone of the instrument's exchange.
	//  - GOOD_TILL_CANCEL: The order remains active indefinitely until
	//    it is either filled or explicitly cancelled by you.
	// Enum: "DAY", "GOOD_TILL_CANCEL"
	TimeInForce string `json:"timeInForce"`
	// Enum: "LIMIT", "STOP", "MARKET", "STOP_LIMIT"
	Type string `json:"type"`
	// The total monetary value of the order. Applicable to value orders.
	Value float64 `json:"value"`
}

type baseOrderRequest struct {
	// Unique instrument identifier. Get from the /instruments endpoint
	Ticker string `json:"ticker"`
	// The total number of shares requested. Applicable to quantity orders.
	Quantity float64 `json:"quantity"`
}
type baseLimitOrderRequest struct {
	LimitPrice  float64 `json:"limitPrice"`
	TimeInForce string  `json:"timeInForce,omitempty"`
}

type baseStopOrderRequest struct {
	StopPrice float64 `json:"stopPrice"`
}

type LimitOrderRequest struct {
	baseOrderRequest
	baseLimitOrderRequest
}

type MarketOrderRequest struct {
	baseOrderRequest

	ExtendedHours bool `json:"extendedHours,omitempty"`
}

type StopLimitOrderRequest struct {
	baseOrderRequest
	baseLimitOrderRequest
}

type StopOrderRequest struct {
	baseOrderRequest
	baseStopOrderRequest
}
