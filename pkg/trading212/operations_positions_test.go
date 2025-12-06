// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"iter"
	"testing"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
)

func Test_Positions_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Positions GetAllPositions", func(t *testing.T) {
			validateOperationItems[models.Position](
				t,
				func(api *API) (iter.Seq[*models.Position], error) {
					return api.Positions.GetAllPositions()
				}, `
			[{
				"averagePricePaid": 0,
				"createdAt": "2019-08-24T14:15:22Z",
				"currentPrice": 0,
				"instrument": {
					"currency": "string",
					"isin": "string",
					"name": "string",
					"ticker": "AAPL_US_EQ"
				},
				"quantity": 0,
				"quantityAvailableForTrading": 0,
				"quantityInPies": 0,
				"walletImpact": {
					"currency": "string",
					"currentValue": 0,
					"fxImpact": 0,
					"totalCost": 0,
					"unrealizedProfitLoss": 0
				}
			}]
		`,
			)
		},
	)
}
