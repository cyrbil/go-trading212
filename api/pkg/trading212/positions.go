// Package trading212 github.com/cyrbil/go-trading212
package trading212

import "time"

// Position response type
type Position struct {
	AveragePricePaid float64   `json:"averagePricePaid"`
	CreatedAt        time.Time `json:"createdAt"`
	CurrentPrice     float64   `json:"currentPrice"`
	Instrument       struct {
		Currency string `json:"currency"`
		Isin     string `json:"isin"`
		Name     string `json:"name"`
		Ticker   string `json:"ticker"`
	} `json:"instrument"`
	Quantity                    float64 `json:"quantity"`
	QuantityAvailableForTrading float64 `json:"quantityAvailableForTrading"`
	QuantityInPies              float64 `json:"quantityInPies"`
	WalletImpact                struct {
		Currency             string  `json:"currency"`
		CurrentValue         float64 `json:"currentValue"`
		FxImpact             float64 `json:"fxImpact"`
		TotalCost            float64 `json:"totalCost"`
		UnrealizedProfitLoss float64 `json:"unrealizedProfitLoss"`
	} `json:"walletImpact"`
}
