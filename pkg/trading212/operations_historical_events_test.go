package trading212

import (
	models "github.com/cyrbil/go-trading212/pkg/trading212/models"
	"iter"
	"testing"
)

func Test_HistoricalEvents_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test HistoricalEvents GetPaidOutDividends", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.Dividend](
				t,
				func(api *API) (iter.Seq[*models.Dividend], error) {
					return api.HistoricalEvents.GetPaidOutDividends()
				}, `
			{
				"items": [{
					"amount": 0,
					"amountInEuro": 0,
					"currency": "string",
					"grossAmountPerShare": 0,
					"instrument": {
						"currency": "string",
						"isin": "string",
						"name": "string",
						"ticker": "AAPL_US_EQ"
					},
					"paidOn": "2019-08-24T14:15:22Z",
					"quantity": 0,
					"reference": "string",
					"ticker": "string",
					"tickerCurrency": "string",
					"type": "ORDINARY"
				}],
				"nextPagePath": null
			}
		`,
			)
		},
	)

	t.Run(
		"Test HistoricalEvents GetHistoricalOrders", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.OrderFill](
				t,
				func(api *API) (iter.Seq[*models.OrderFill], error) {
					return api.HistoricalEvents.GetHistoricalOrders()
				}, `
			{
				"items": [{
					"fill": {
						"filledAt": "2019-08-24T14:15:22Z",
						"id": 0,
						"price": 0,
						"quantity": 0,
						"tradingMethod": "TOTV",
						"type": "TRADE",
						"walletImpact": {
							"currency": "string",
							"fxRate": 0,
							"netValue": 0,
							"realisedProfitLoss": 0,
							"taxes": [{
								"chargedAt": "2019-08-24T14:15:22Z",
								"currency": "string",
								"name": "COMMISSION_TURNOVER",
								"quantity": 0
							}]
						}
					},
					"order": {
						"createdAt": "2019-08-24T14:15:22Z",
						"currency": "string",
						"extendedHours": true,
						"filledQuantity": 0,
						"filledValue": 0,
						"id": 0,
						"initiatedFrom": "API",
						"instrument": {
							"currency": "string",
							"isin": "string",
							"name": "string",
							"ticker": "AAPL_US_EQ"
						},
						"limitPrice": 0,
						"quantity": 0,
						"side": "BUY",
						"status": "LOCAL",
						"stopPrice": 0,
						"strategy": "QUANTITY",
						"ticker": "AAPL_US_EQ",
						"timeInForce": "DAY",
						"type": "LIMIT",
						"value": 0
					}
				}],
				"nextPagePath": null
			}
		`,
			)
		},
	)

	t.Run(
		"Test HistoricalEvents GetTransactions", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.Transaction](
				t,
				func(api *API) (iter.Seq[*models.Transaction], error) {
					return api.HistoricalEvents.GetTransactions()
				}, `
			{
				"items": [{
					"amount": 0,
					"currency": "string",
					"dateTime": "2019-08-24T14:15:22Z",
					"reference": "string",
					"type": "WITHDRAW"
				}],
				"nextPagePath": null
			}
		`,
			)
		},
	)

	t.Run(
		"Test HistoricalEvents ListReports", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[models.Report](
				t,
				func(api *API) (iter.Seq[*models.Report], error) {
					return api.HistoricalEvents.ListReports()
				}, `
			[{
				"dataIncluded": {
					"includeDividends": true,
					"includeInterest": true,
					"includeOrders": true,
					"includeTransactions": true
				},
				"downloadLink": "string",
				"reportId": 0,
				"status": "Queued",
				"timeFrom": "2019-08-24T14:15:22Z",
				"timeTo": "2019-08-24T14:15:22Z"
			}]
		`,
			)
		},
	)

	t.Run(
		"Test HistoricalEvents RequestReport", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[models.ReportID](
				t,
				func(api *API) (*models.ReportID, error) {
					return api.HistoricalEvents.RequestReport(models.ReportRequest{})
				}, `
			{
				"reportId": 0
			}
		`,
			)
		},
	)
}
