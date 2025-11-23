package trading212

import "time"

type PieSummary struct {
	Cash            float64 `json:"cash"`
	DividendDetails struct {
		Gained     float64 `json:"gained"`
		InCash     float64 `json:"inCash"`
		Reinvested float64 `json:"reinvested"`
	} `json:"dividendDetails"`
	Id       uint    `json:"id"`
	Progress float64 `json:"progress"`
	Result   struct {
		PriceAvgInvestedValue float64 `json:"priceAvgInvestedValue"`
		PriceAvgResult        float64 `json:"priceAvgResult"`
		PriceAvgResultCoef    float64 `json:"priceAvgResultCoef"`
		PriceAvgValue         float64 `json:"priceAvgValue"`
	} `json:"result"`
	Status string `json:"status"`
}

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
		CreationDate       time.Time          `json:"creationDate"`
		DividendCashAction string             `json:"dividendCashAction"`
		EndDate            time.Time          `json:"endDate"`
		Goal               float64            `json:"goal"`
		Icon               string             `json:"icon"`
		Id                 uint               `json:"id"`
		InitialInvestment  float64            `json:"initialInvestment"`
		InstrumentShares   map[string]float64 `json:"instrumentShares"`
		Name               string             `json:"name"`
		PublicUrl          string             `json:"publicUrl"`
	} `json:"settings"`
}

type PieRequest struct {
	PieMetaRequest

	DividendCashAction string    `json:"dividendCashAction"`
	EndDate            time.Time `json:"endDate"`
	Goal               float64   `json:"goal"`
	InstrumentShares   struct {
		AAPLUSEQ float64 `json:"AAPL_US_EQ"`
		MSFTUSEQ float64 `json:"MSFT_US_EQ"`
	} `json:"instrumentShares"`
}

type PieMetaRequest struct {
	Icon string `json:"icon"`
	Name string `json:"name"`
}
