package models

import "time"

// PieSummary response type.
type PieSummary struct {
	// Cash.
	Cash float64 `json:"cash"`
	// DividendDetails.
	DividendDetails struct {
		// Gained.
		Gained float64 `json:"gained"`
		// InCash.
		InCash float64 `json:"inCash"`
		// Reinvested.
		Reinvested float64 `json:"reinvested"`
	} `json:"dividendDetails"`
	// ID.
	ID uint `json:"id"`
	// Progress.
	Progress float64 `json:"progress"`
	// Result.
	Result struct {
		// PriceAvgInvestedValue.
		PriceAvgInvestedValue float64 `json:"priceAvgInvestedValue"`
		// PriceAvgResult.
		PriceAvgResult float64 `json:"priceAvgResult"`
		// PriceAvgResultCoef.
		PriceAvgResultCoef float64 `json:"priceAvgResultCoef"`
		// PriceAvgValue.
		PriceAvgValue float64 `json:"priceAvgValue"`
	} `json:"result"`
	// Status.
	Status string `json:"status"`
}

// PieDetails response type.
type PieDetails struct {
	Instruments []struct {
		CurrentShare  float64 `json:"currentShare"`
		ExpectedShare float64 `json:"expectedShare"`
		Issues        []struct {
			Name     string `json:"name"`
			Severity string `json:"severity"`
		} `json:"issues"`
		OwnedQuantity float64 `json:"ownedQuantity"`
		Result        struct {
			PriceAvgInvestedValue float64 `json:"priceAvgInvestedValue"`
			PriceAvgResult        float64 `json:"priceAvgResult"`
			PriceAvgResultCoef    float64 `json:"priceAvgResultCoef"`
			PriceAvgValue         float64 `json:"priceAvgValue"`
		} `json:"result"`
		Ticker string `json:"ticker"`
	} `json:"instruments"`
	Settings struct {
		ID                 uint               `json:"id"`
		CreationDate       time.Time          `json:"creationDate"`
		DividendCashAction string             `json:"dividendCashAction"`
		EndDate            time.Time          `json:"endDate"`
		Goal               float64            `json:"goal"`
		Icon               string             `json:"icon"`
		InitialInvestment  float64            `json:"initialInvestment"`
		InstrumentShares   map[string]float64 `json:"instrumentShares"`
		Name               string             `json:"name"`
		PublicURL          string             `json:"publicUrl"`
	} `json:"settings"`
}

// PieRequest request type.
type PieRequest struct {
	PieMetaRequest

	DividendCashAction string    `json:"dividendCashAction"`
	EndDate            time.Time `json:"endDate"`
	Goal               float64   `json:"goal"`
	InstrumentShares   struct {
		AAPLUSEQ float64 `json:"AAPL_US_EQ"` //nolint:tagliatelle
		MSFTUSEQ float64 `json:"MSFT_US_EQ"` //nolint:tagliatelle
	} `json:"instrumentShares"`
}

// PieMetaRequest request type.
type PieMetaRequest struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
}
