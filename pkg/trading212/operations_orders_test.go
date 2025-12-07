package trading212

import (
	trading213 "github.com/cyrbil/go-trading212/pkg/trading212/models"
	"iter"
	"testing"
)

func Test_Orders_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Orders GetAllPendingOrders", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[trading213.Order](
				t,
				func(api *API) (iter.Seq[*trading213.Order], error) {
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
			t.Parallel()
			validateOperationObject[trading213.Order](
				t,
				func(api *API) (*trading213.Order, error) {
					return api.Orders.PlaceLimitOrder(trading213.LimitOrderRequest{})
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
			t.Parallel()
			validateOperationObject[trading213.Order](
				t,
				func(api *API) (*trading213.Order, error) {
					return api.Orders.PlaceMarketOrder(trading213.MarketOrderRequest{})
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
			t.Parallel()
			validateOperationObject[trading213.Order](
				t,
				func(api *API) (*trading213.Order, error) {
					return api.Orders.PlaceStopOrder(trading213.StopOrderRequest{})
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
			t.Parallel()
			validateOperationObject[trading213.Order](
				t,
				func(api *API) (*trading213.Order, error) {
					return api.Orders.PlaceStopLimitOrder(trading213.StopLimitOrderRequest{})
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
			t.Parallel()
			validateOperationObject[trading213.Empty](
				t,
				func(api *API) (*trading213.Empty, error) {
					return &trading213.Empty{}, api.Orders.CancelOrder(0)
				}, "",
			)
		},
	)

	t.Run(
		"Test Orders GetPendingOrderByID", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.Order](
				t,
				func(api *API) (*trading213.Order, error) {
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
