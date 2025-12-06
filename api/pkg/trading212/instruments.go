// Package trading212 github.com/cyrbil/go-trading212
package trading212

import "time"

// ExchangeMetadata response type
type ExchangeMetadata struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	WorkingSchedules []struct {
		ID         uint `json:"id"`
		TimeEvents []struct {
			Date time.Time `json:"date"`
			//  Enum "OPEN", "CLOSE", "BREAK_START", "BREAK_END", "PRE_MARKET_OPEN",
			//  "AFTER_HOURS_OPEN", "AFTER_HOURS_CLOSE", "OVERNIGHT_OPEN"
			Type string `json:"type"`
		} `json:"timeEvents"`
	} `json:"workingSchedules"`
}

// Instrument response type
type Instrument struct {
	// On the platform since
	AddedOn time.Time `json:"addedOn"`
	// ISO 4217
	CurrencyCode    string  `json:"currencyCode"`
	ExtendedHours   bool    `json:"extendedHours"`
	Isin            string  `json:"isin"`
	MaxOpenQuantity float64 `json:"maxOpenQuantity"`
	Name            string  `json:"name"`
	ShortName       string  `json:"shortName"`
	// Unique identifier
	Ticker string `json:"ticker"`
	// Enum "CRYPTOCURRENCY", "ETF", "FOREX", "FUTURES", "INDEX", "STOCK", "WARRANT", "CRYPTO", "CVR", "CORPACT"
	Type string `json:"type"`
	// Get items in the /exchanges endpoint
	WorkingScheduleID uint `json:"workingScheduleId"`
}
