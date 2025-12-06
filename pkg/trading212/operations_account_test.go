package trading212

import (
	"testing"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
)

func Test_Account_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Account.GetAccountSummary()", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[models.AccountSummary](
				t,
				func(api *API) (*models.AccountSummary, error) {
					return api.Account.GetAccountSummary()
				}, `
			{
			  "cash": {
				"availableToTrade": 0,
				"inPies": 0,
				"reservedForOrders": 0
			  },
			  "currency": "string",
			  "id": 0,
			  "investments": {
				"currentValue": 0,
				"realizedProfitLoss": 0,
				"totalCost": 0,
				"unrealizedProfitLoss": 0
			  },
			  "totalValue": 0
			}
		`,
			)
		},
	)
}
