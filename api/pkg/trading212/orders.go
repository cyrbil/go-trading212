package trading212

import "time"

type Order struct {
	CreatedAt      time.Time `json:"createdAt"`
	Currency       string    `json:"currency"`
	ExtendedHours  bool      `json:"extendedHours"`
	FilledQuantity float64   `json:"filledQuantity"`
	FilledValue    float64   `json:"filledValue"`
	Id             uint      `json:"id"`
	InitiatedFrom  string    `json:"initiatedFrom"`
	Instrument     struct {
		Currency string `json:"currency"`
		Isin     string `json:"isin"`
		Name     string `json:"name"`
		Ticker   string `json:"ticker"`
	} `json:"instrument"`
	LimitPrice  float64 `json:"limitPrice"`
	Quantity    float64 `json:"quantity"`
	Side        string  `json:"side"`
	Status      string  `json:"status"`
	StopPrice   float64 `json:"stopPrice"`
	Strategy    string  `json:"strategy"`
	Ticker      string  `json:"ticker"`
	TimeInForce string  `json:"timeInForce"`
	Type        string  `json:"type"`
	Value       float64 `json:"value"`
}

type baseOrderRequest struct {
	Ticker   string  `json:"ticker"`
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
