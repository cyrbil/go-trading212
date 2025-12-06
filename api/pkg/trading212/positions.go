package trading212

import "time"

// Position response type.
type Position struct {
	// AveragePricePaid.
	AveragePricePaid float64 `json:"averagePricePaid"`
	// CreatedAt.
	CreatedAt time.Time `json:"createdAt"`
	// CurrentPrice.
	CurrentPrice float64 `json:"currentPrice"`
	// Instrument.
	Instrument struct {
		// Currency.
		Currency string `json:"currency"`
		// Isin.
		Isin string `json:"isin"`
		// Name.
		Name string `json:"name"`
		// Ticker.
		Ticker string `json:"ticker"`
	} `json:"instrument"`
	// Quantity.
	Quantity float64 `json:"quantity"`
	// QuantityAvailableForTrading.
	QuantityAvailableForTrading float64 `json:"quantityAvailableForTrading"`
	// QuantityInPies.
	QuantityInPies float64 `json:"quantityInPies"`
	// WalletImpact.
	WalletImpact struct {
		// Currency.
		Currency string `json:"currency"`
		// CurrentValue.
		CurrentValue float64 `json:"currentValue"`
		// FxImpact.
		FxImpact float64 `json:"fxImpact"`
		// TotalCost.
		TotalCost float64 `json:"totalCost"`
		// UnrealizedProfitLoss.
		UnrealizedProfitLoss float64 `json:"unrealizedProfitLoss"`
	} `json:"walletImpact"`
}
