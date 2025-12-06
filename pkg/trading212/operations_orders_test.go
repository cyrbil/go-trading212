// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"iter"
	"testing"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
)

func Test_Orders_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Orders GetAllPendingOrders", func(t *testing.T) {
			validateOperationItems[models.Order](
				t,
				func(api *API) (iter.Seq[*models.Order], error) {
					return api.Orders.GetAllPendingOrders()
				}, `
			[{
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
			}]
		`,
			)
		},
	)

	t.Run(
		"Test Orders PlaceLimitOrder", func(t *testing.T) {
			validateOperationObject[models.Order](
				t,
				func(api *API) (*models.Order, error) {
					return api.Orders.PlaceLimitOrder(models.LimitOrderRequest{})
				}, `
			{
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
		`,
			)
		},
	)

	t.Run(
		"Test Orders PlaceMarketOrder", func(t *testing.T) {
			validateOperationObject[models.Order](
				t,
				func(api *API) (*models.Order, error) {
					return api.Orders.PlaceMarketOrder(models.MarketOrderRequest{})
				}, `
			{
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
		`,
			)
		},
	)

	t.Run(
		"Test Orders PlaceStopOrder", func(t *testing.T) {
			validateOperationObject[models.Order](
				t,
				func(api *API) (*models.Order, error) {
					return api.Orders.PlaceStopOrder(models.StopOrderRequest{})
				}, `
			{
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
		`,
			)
		},
	)

	t.Run(
		"Test Orders PlaceStopLimitOrder", func(t *testing.T) {
			validateOperationObject[models.Order](
				t,
				func(api *API) (*models.Order, error) {
					return api.Orders.PlaceStopLimitOrder(models.StopLimitOrderRequest{})
				}, `
			{
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
		`,
			)
		},
	)

	t.Run(
		"Test Orders CancelOrder", func(t *testing.T) {
			validateOperationObject[models.Empty](
				t,
				func(api *API) (*models.Empty, error) {
					return &models.Empty{}, api.Orders.CancelOrder(0)
				}, "",
			)
		},
	)

	t.Run(
		"Test Orders GetPendingOrderByID", func(t *testing.T) {
			validateOperationObject[models.Order](
				t,
				func(api *API) (*models.Order, error) {
					return api.Orders.GetPendingOrderByID(0)
				}, `
			{
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
		`,
			)
		},
	)
}
