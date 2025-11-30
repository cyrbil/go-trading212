package trading212

import (
	"fmt"
	"iter"
	"net/http"
)

import (
	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type operationGetAllPendingOrders interface {
	// GetAllPendingOrders operation.
	// Retrieves a list of all orders that are currently active (i.e., not yet filled, cancelled, or expired).
	// This is useful for monitoring the status of your open positions and managing your trading strategy.
	// See: https://docs.trading212.com/api/orders/orders
	GetAllPendingOrders() (iter.Seq[models.Order], error)
}

type operationPlaceLimitOrder interface {
	// PlaceLimitOrder operation.
	// Creates a new Limit order, which executes at a specified price or better.
	// To place a buy order, use a positive quantity. The order will fill at the limitPrice or lower.
	// To place a sell order, use a negative quantity. The order will fill at the limitPrice or higher.
	// See: https://docs.trading212.com/api/orders/placelimitorder
	PlaceLimitOrder(req *models.LimitOrderRequest) (*models.Order, error)
}

type operationPlaceMarketOrder interface {
	// PlaceMarketOrder operation.
	// Creates a new Market order, which is an instruction to trade a security immediately at the next available price.
	// To place a buy order, use a positive quantity.
	// To place a sell order, use a negative quantity.
	// extendedHours: Set to true to allow the order to be filled outside the standard trading session.
	// If placed when the market is closed, the order will be queued to execute when the market next opens.
	// See: https://docs.trading212.com/api/orders/placemarketorder
	PlaceMarketOrder(req *models.MarketOrderRequest) (*models.Order, error)
}

type operationPlaceStopOrder interface {
	// PlaceStopOrder operation.
	// Creates a new Stop order, which places a Market order once the stopPrice is reached.
	// To place a buy stop order, use a positive quantity.
	// To place a sell stop order (commonly a 'stop-loss'), use a negative quantity.
	// The stopPrice is triggered by the instrument's Last Traded Price (LTP).
	// See: https://docs.trading212.com/api/orders/placestoporder_1
	PlaceStopOrder(req *models.StopOrderRequest) (*models.Order, error)
}

type operationPlaceStopLimitOrder interface {
	// PlaceStopLimitOrder operation.
	// Creates a new Stop-Limit order, combining features of a Stop and a Limit order.
	// The direction of the trade (buy/sell) is determined by the sign of the quantity field.
	// Execution Logic:
	//   - When the instrument's Last Traded Price (LTP) reaches the specified stopPrice, the order is triggered.
	//   - A Limit order is then automatically placed at the specified limitPrice.
	//
	// This two-step process helps protect against price slippage that can occur with a standard Stop order.
	// See: https://docs.trading212.com/api/orders/placestoporder
	PlaceStopLimitOrder(req models.StopLimitOrderRequest) (*models.Order, error)
}

type operationCancelOrder interface {
	// CancelOrder operation.
	// Attempts to cancel an active, unfilled order by its unique ID. Cancellation
	// is not guaranteed if the order is already in the process of being filled. A
	// successful response indicates the cancellation request was accepted.
	// See: https://docs.trading212.com/api/orders/cancelorder
	CancelOrder(id int64) (*models.Empty, error)
}

type operationGetPendingOrderByID interface {
	// GetPendingOrderByID operation.
	// Retrieves a single pending order using its unique numerical ID.
	// This is useful for checking the status of a specific order you have previously placed.
	// See: https://docs.trading212.com/api/orders/orderbyid
	GetPendingOrderByID(id int64) (*models.Order, error)
}

type ordersOperations interface {
	operationGetAllPendingOrders
	operationPlaceLimitOrder
	operationPlaceMarketOrder
	operationPlaceStopOrder
	operationPlaceStopLimitOrder
	operationCancelOrder
	operationGetPendingOrderByID
}

type orders struct {
	api requestMaker
}

func (op *orders) GetAllPendingOrders() (iter.Seq[models.Order], error) {
	return runOperation[models.Order](op.api, http.MethodGet, internal.GetAllPendingOrders, nil).Items()
}

func (op *orders) PlaceLimitOrder(req *models.LimitOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceLimitOrder, req).Object()
}

func (op *orders) PlaceMarketOrder(req *models.MarketOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceMarketOrder, jsonBody{req}).Object()
}

func (op *orders) PlaceStopOrder(req *models.StopOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceStopOrder, jsonBody{req}).Object()
}

func (op *orders) PlaceStopLimitOrder(req models.StopLimitOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceStopLimitOrder, jsonBody{req}).Object()
}

func (op *orders) CancelOrder(id int64) (*models.Empty, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.CancelOrder, id))
	return runOperation[models.Empty](op.api, http.MethodDelete, endpoint, nil).Object()
}

func (op *orders) GetPendingOrderByID(id int64) (*models.Order, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.GetPendingOrderByID, id))
	return runOperation[models.Order](op.api, http.MethodGet, endpoint, nil).Object()
}
