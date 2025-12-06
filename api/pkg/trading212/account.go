package trading212

// AccountSummary response type.
type AccountSummary struct {
	// Primary trading account number.
	// This is the same account ID you would see in the Trading 212 web or mobile application.
	ID uint `json:"id"`
	// Primary account currency in ISO 4217 format.
	Currency string `json:"currency"`
	// Investments value in your account's primary currency.
	TotalValue float64 `json:"totalValue"`
	Cash       struct {
		// Funds available for investing.
		AvailableToTrade float64 `json:"availableToTrade"`
		// The amount of cash reserved for pending orders. This cash is not available for placing new trades.
		ReservedForOrders float64 `json:"reservedForOrders"`
		// The amount of cash in Pies.
		InPies float64 `json:"inPies"`
	} `json:"cash"`
	Investments struct {
		// Current value of all the investments.
		CurrentValue float64 `json:"currentValue"`
		// The all-time realised profit loss from all the trades executed.
		TotalCost float64 `json:"totalCost"`
		// The cost basis of your current investments.
		// The total amount of funds you've invested in the shares you currently own.
		RealizedProfitLoss float64 `json:"realizedProfitLoss"`
		// The potential profit/loss of your current investments,
		// showing how much you could gain or lose if you were to sell them now.
		UnrealizedProfitLoss float64 `json:"unrealizedProfitLoss"`
	} `json:"investments"`
}
