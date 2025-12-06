package trading212

import (
	"iter"
	"testing"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
)

func Test_Instruments_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Instruments GetExchangesMetadata", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.ExchangeMetadata](
				t,
				func(api *API) (iter.Seq[*models.ExchangeMetadata], error) {
					return api.Instruments.GetExchangesMetadata()
				}, `
			[{
				"id": 0,
				"name": "string",
				"workingSchedules": [{
					"id": 0,
					"timeEvents": [{
						"date": "2019-08-24T14:15:22Z",
						"type": "OPEN"
					}]
				}]
			}]
		`,
			)
		},
	)

	t.Run(
		"Test Instruments GetAllAvailableInstruments", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.Instrument](
				t,
				func(api *API) (iter.Seq[*models.Instrument], error) {
					return api.Instruments.GetAllAvailableInstruments()
				}, `
			[{
				"addedOn": "2019-08-24T14:15:22Z",
				"currencyCode": "USD",
				"extendedHours": true,
				"isin": "string",
				"maxOpenQuantity": 0,
				"name": "string",
				"shortName": "string",
				"ticker": "AAPL_US_EQ",
				"type": "ETF",
				"workingScheduleId": 0
			}]
		`,
			)
		},
	)
}
