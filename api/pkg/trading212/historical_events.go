package trading212

import "time"

type Dividend struct {
	Ticker              string    `json:"ticker"`
	Reference           string    `json:"reference"`
	Quantity            float64   `json:"quantity"`
	Amount              float64   `json:"amount"`
	GrossAmountPerShare float64   `json:"grossAmountPerShare"`
	AmountInEuro        float64   `json:"amountInEuro"`
	PaidOn              time.Time `json:"paidOn"`
	Type                string    `json:"type"`
}

type ReportId struct {
	ReportId uint `json:"reportId"`
}

type Report struct {
	ReportId

	DataIncluded struct {
		IncludeDividends    bool `json:"includeDividends"`
		IncludeInterest     bool `json:"includeInterest"`
		IncludeOrders       bool `json:"includeOrders"`
		IncludeTransactions bool `json:"includeTransactions"`
	} `json:"dataIncluded"`
	DownloadLink string    `json:"downloadLink"`
	Status       string    `json:"status"`
	TimeFrom     time.Time `json:"timeFrom"`
	TimeTo       time.Time `json:"timeTo"`
}

type ReportRequest struct {
	DataIncluded struct {
		IncludeDividends    bool `json:"includeDividends"`
		IncludeInterest     bool `json:"includeInterest"`
		IncludeOrders       bool `json:"includeOrders"`
		IncludeTransactions bool `json:"includeTransactions"`
	} `json:"dataIncluded"`
	TimeFrom time.Time `json:"timeFrom"`
	TimeTo   time.Time `json:"timeTo"`
}

type OrderFill struct {
	Order `json:"order"`

	Fill struct {
		FilledAt      time.Time `json:"filledAt"`
		Id            uint      `json:"id"`
		Price         float64   `json:"price"`
		Quantity      float64   `json:"quantity"`
		TradingMethod string    `json:"tradingMethod"`
		Type          string    `json:"type"`
		WalletImpact  struct {
			Currency             string `json:"currency"`
			CurrentValue         int    `json:"currentValue"`
			FxImpact             int    `json:"fxImpact"`
			TotalCost            int    `json:"totalCost"`
			UnrealizedProfitLoss int    `json:"unrealizedProfitLoss"`
		} `json:"walletImpact"`
	} `json:"fill"`
}

type Transaction struct {
	Amount    uint      `json:"amount"`
	DateTime  time.Time `json:"dateTime"`
	Reference string    `json:"reference"`
	Type      string    `json:"type"`
}
