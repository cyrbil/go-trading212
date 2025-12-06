// Package trading212 github.com/cyrbil/go-trading212
package trading212

import "time"

// Dividend response type
type Dividend struct {
	// Amount, in account's primary currency.
	Amount int `json:"amount"`
	// AmountInEuro
	AmountInEuro int `json:"amountInEuro"`
	// Currency, the account's primary currency.
	Currency string `json:"currency"`
	// GrossAmountPerShare, in instrument currency
	GrossAmountPerShare int `json:"grossAmountPerShare"`
	// Instrument information as given by /instruments endpoint.
	Instrument struct {
		// Currency, instrument currency in ISO 4217 format.
		Currency string `json:"currency"`
		// ISIN of the instrument.
		ISIN string `json:"isin"`
		// Name or the instrument.
		Name string `json:"name"`
		// Ticker, unique instrument identifier.
		Ticker string `json:"ticker"`
	} `json:"instrument"`
	// PaidOn
	PaidOn time.Time `json:"paidOn"`
	// Quantity
	Quantity int `json:"quantity"`
	// Reference
	Reference string `json:"reference"`
	// Ticker
	Ticker string `json:"ticker"`
	// TickerCurrency
	TickerCurrency string `json:"tickerCurrency"`
	// Type, enum "ORDINARY", "BONUS", "PROPERTY_INCOME", "RETURN_OF_CAPITAL_NON_US", "DEMERGER", "INTEREST",
	// "CAPITAL_GAINS_DISTRIBUTION_NON_US", "INTERIM_LIQUIDATION", "ORDINARY_MANUFACTURED_PAYMENT",
	// "BONUS_MANUFACTURED_PAYMENT", "PROPERTY_INCOME_MANUFACTURED_PAYMENT",
	// "RETURN_OF_CAPITAL_NON_US_MANUFACTURED_PAYMENT", "DEMERGER_MANUFACTURED_PAYMENT",
	// "INTEREST_MANUFACTURED_PAYMENT", "CAPITAL_GAINS_DISTRIBUTION_NON_US_MANUFACTURED_PAYMENT",
	// "INTERIM_LIQUIDATION_MANUFACTURED_PAYMENT", "INTEREST_PAID_BY_US_OBLIGORS",
	// "INTEREST_PAID_BY_FOREIGN_CORPORATIONS", "DIVIDENDS_PAID_BY_US_CORPORATIONS",
	// "DIVIDENDS_PAID_BY_FOREIGN_CORPORATIONS", "CAPITAL_GAINS",
	// "REAL_PROPERTY_INCOME_AND_NATURAL_RESOURCES_ROYALTIES", "OTHER_INCOME", "QUALIFIED_INVESTMENT_ENTITY",
	// "TRUST_DISTRIBUTION", "PUBLICLY_TRADED_PARTNERSHIP_DISTRIBUTION", "CAPITAL_GAINS_DISTRIBUTION",
	// "RETURN_OF_CAPITAL", "OTHER_DIVIDEND_EQUIVALENT", "TAX_EVENT_1446F_FOR_PUBLICLY_TRADED_SECURITIES",
	// "PTP_UNCHARACTERISED_INCOME", "MULTIPLE_1042S_TAX_COMPONENTS", "DIVIDEND", "SHORT_TERM_CAPITAL_GAINS",
	// "LONG_TERM_CAPITAL_GAINS", "PROPERTY_INCOME_DISTRIBUTION", "TAX_EXEMPTED",
	// "INTEREST_PAID_BY_US_OBLIGORS_MANUFACTURED_PAYMENT",
	// "INTEREST_PAID_BY_FOREIGN_CORPORATIONS_MANUFACTURED_PAYMENT",
	// "DIVIDENDS_PAID_BY_US_CORPORATIONS_MANUFACTURED_PAYMENT",
	// "DIVIDENDS_PAID_BY_FOREIGN_CORPORATIONS_MANUFACTURED_PAYMENT",
	// "CAPITAL_GAINS_MANUFACTURED_PAYMENT",
	// "REAL_PROPERTY_INCOME_AND_NATURAL_RESOURCES_ROYALTIES_MANUFACTURED_PAYMENT", "OTHER_INCOME_MANUFACTURED_PAYMENT",
	// "QUALIFIED_INVESTMENT_ENTITY_MANUFACTURED_PAYMENT", "TRUST_DISTRIBUTION_MANUFACTURED_PAYMENT",
	// "PUBLICLY_TRADED_PARTNERSHIP_DISTRIBUTION_MANUFACTURED_PAYMENT",
	// "CAPITAL_GAINS_DISTRIBUTION_MANUFACTURED_PAYMENT", "RETURN_OF_CAPITAL_MANUFACTURED_PAYMENT",
	// "OTHER_DIVIDEND_EQUIVALENT_MANUFACTURED_PAYMENT",
	// "TAX_EVENT_1446F_FOR_PUBLICLY_TRADED_SECURITIES_MANUFACTURED_PAYMENT",
	// "PTP_UNCHARACTERISED_INCOME_MANUFACTURED_PAYMENT", "MULTIPLE_1042S_TAX_COMPONENTS_MANUFACTURED_PAYMENT",
	// "DIVIDEND_MANUFACTURED_PAYMENT", "SHORT_TERM_CAPITAL_GAINS_MANUFACTURED_PAYMENT",
	// "LONG_TERM_CAPITAL_GAINS_MANUFACTURED_PAYMENT", "PROPERTY_INCOME_DISTRIBUTION_MANUFACTURED_PAYMENT",
	// "TAX_EXEMPTED_MANUFACTURED_PAYMENT"
	Type string `json:"type"`
}

// ReportID response type
type ReportID struct {
	ReportID uint `json:"reportId"`
}

type reportDataIncluded struct {
	DataIncluded struct {
		IncludeDividends    bool `json:"includeDividends"`
		IncludeInterest     bool `json:"includeInterest"`
		IncludeOrders       bool `json:"includeOrders"`
		IncludeTransactions bool `json:"includeTransactions"`
	} `json:"dataIncluded"`
}

type reportTimeRange struct {
	TimeFrom time.Time `json:"timeFrom"`
	TimeTo   time.Time `json:"timeTo"`
}

// Report response type
type Report struct {
	ReportID
	reportDataIncluded
	reportTimeRange

	DownloadLink string `json:"downloadLink"`
	Status       string `json:"status"`
}

// ReportRequest request type
type ReportRequest struct {
	reportDataIncluded
	reportTimeRange
}

// OrderFill response type
type OrderFill struct {
	Order `json:"order"`

	Fill struct {
		FilledAt      time.Time `json:"filledAt"`
		ID            int       `json:"id"`
		Price         int       `json:"price"`
		Quantity      int       `json:"quantity"`
		TradingMethod string    `json:"tradingMethod"`
		Type          string    `json:"type"`
		WalletImpact  struct {
			Currency           string `json:"currency"`
			FxRate             int    `json:"fxRate"`
			NetValue           int    `json:"netValue"`
			RealisedProfitLoss int    `json:"realisedProfitLoss"`
			Taxes              []struct {
				ChargedAt time.Time `json:"chargedAt"`
				Currency  string    `json:"currency"`
				Name      string    `json:"name"`
				Quantity  int       `json:"quantity"`
			} `json:"taxes"`
		} `json:"walletImpact"`
	} `json:"fill"`
}

// Transaction response type
type Transaction struct {
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	DateTime  time.Time `json:"dateTime"`
	Reference string    `json:"reference"`
	Type      string    `json:"type"`
}
