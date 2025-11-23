package trading212

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"iter"
	"net/http"
)

import (
	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

// helper struct to have a json reader object
type jsonBody struct{ data any }

func (b jsonBody) Read(p []byte) (n int, err error) {
	jsonData, err := json.Marshal(b.data)
	if err != nil {
		return 0, err
	}
	return bytes.NewReader(jsonData).Read(p)
}

// helper function for the operations
func runOperation[T any](api *API, method string, endpoint internal.APIEndpoint, body any) *Response[T] {
	var requestBodyReader io.Reader
	if body != nil {
		requestBodyReader = jsonBody{body}
	}
	request, err := api.NewRequest(method, endpoint, requestBodyReader)
	if err != nil {
		return &Response[T]{err: err}
	}
	data, err := request.Do()
	if err != nil {
		return &Response[T]{err: err}
	}

	return &Response[T]{request: request, raw: data}
}

// Account operations.
// Access fundamental information about your trading account.
// Retrieve details such as your account ID, currency, and current cash balance.
// See: https://docs.trading212.com/api/accounts
func (api *API) Account() *Account {
	return &Account{api}
}

type Account struct {
	api *API
}

// GetAccountSummary operation.
// Provides a breakdown of your account's cash and investment metrics,
// including available funds, invested capital, and total account value.
// See: https://docs.trading212.com/api/accounts/getaccountsummary
func (op *Account) GetAccountSummary() (*models.AccountSummary, error) {
	return runOperation[models.AccountSummary](op.api, http.MethodGet, internal.GetAccountSummary, nil).Object()
}

// Instruments operations.
// Discover what you can trade. These endpoints provide comprehensive lists of all tradable instruments
// and the exchanges they belong to, including details like tickers and trading hours.
// See: https://docs.trading212.com/api/instruments
func (api *API) Instruments() *Instruments {
	return &Instruments{api}
}

type Instruments struct {
	api *API
}

// GetExchangesMetadata operation.
// Retrieves all accessible exchanges and their corresponding working schedules. Data is refreshed every 10 minutes.
// See: https://docs.trading212.com/api/instruments/exchanges
func (op *Instruments) GetExchangesMetadata() (iter.Seq[models.ExchangeMetadata], error) {
	return runOperation[models.ExchangeMetadata](op.api, http.MethodGet, internal.GetExchangesMetadata, nil).Items()
}

// GetAllAvailableInstruments operation.
// Retrieves all accessible instruments. Data is refreshed every 10 minutes.
// See: https://docs.trading212.com/api/instruments/instruments
func (op *Instruments) GetAllAvailableInstruments() (iter.Seq[models.Instrument], error) {
	return runOperation[models.Instrument](op.api, http.MethodGet, internal.GetAllAvailableInstruments, nil).Items()
}

// Orders operations
// Place, monitor, and cancel equity trade orders.
// This section provides the core functionality for
// programmatically executing your trading strategies for stocks and ETFs.
// See: https://docs.trading212.com/api/orders
func (api *API) Orders() *Orders {
	return &Orders{api}
}

type Orders struct {
	api *API
}

// GetAllPendingOrders operation.
// Retrieves a list of all orders that are currently active (i.e., not yet filled, cancelled, or expired).
// This is useful for monitoring the status of your open positions and managing your trading strategy.
// See: https://docs.trading212.com/api/orders/orders
func (op *Orders) GetAllPendingOrders() (iter.Seq[models.Order], error) {
	return runOperation[models.Order](op.api, http.MethodGet, internal.GetAllPendingOrders, nil).Items()
}

// PlaceLimitOrder operation.
// Creates a new Limit order, which executes at a specified price or better.
// To place a buy order, use a positive quantity. The order will fill at the limitPrice or lower.
// To place a sell order, use a negative quantity. The order will fill at the limitPrice or higher.
// See: https://docs.trading212.com/api/orders/placelimitorder
func (op *Orders) PlaceLimitOrder(req *models.LimitOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceLimitOrder, req).Object()
}

// PlaceMarketOrder operation.
// Creates a new Market order, which is an instruction to trade a security immediately at the next available price.
// To place a buy order, use a positive quantity.
// To place a sell order, use a negative quantity.
// extendedHours: Set to true to allow the order to be filled outside the standard trading session.
// If placed when the market is closed, the order will be queued to execute when the market next opens.
// See: https://docs.trading212.com/api/orders/placemarketorder
func (op *Orders) PlaceMarketOrder(req *models.MarketOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceMarketOrder, jsonBody{req}).Object()
}

// PlaceStopOrder operation.
// Creates a new Stop order, which places a Market order once the stopPrice is reached.
// To place a buy stop order, use a positive quantity.
// To place a sell stop order (commonly a 'stop-loss'), use a negative quantity.
// The stopPrice is triggered by the instrument's Last Traded Price (LTP).
// See: https://docs.trading212.com/api/orders/placestoporder_1
func (op *Orders) PlaceStopOrder(req *models.StopOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceStopOrder, jsonBody{req}).Object()
}

// PlaceStopLimitOrder operation.
// Creates a new Stop-Limit order, combining features of a Stop and a Limit order.
// The direction of the trade (buy/sell) is determined by the sign of the quantity field.
// Execution Logic:
//   - When the instrument's Last Traded Price (LTP) reaches the specified stopPrice, the order is triggered.
//   - A Limit order is then automatically placed at the specified limitPrice.
//
// This two-step process helps protect against price slippage that can occur with a standard Stop order.
// See: https://docs.trading212.com/api/orders/placestoporder
func (op *Orders) PlaceStopLimitOrder(req models.StopLimitOrderRequest) (*models.Order, error) {
	return runOperation[models.Order](op.api, http.MethodPost, internal.PlaceStopLimitOrder, jsonBody{req}).Object()
}

// CancelOrder operation.
// Attempts to cancel an active, unfilled order by its unique ID. Cancellation
// is not guaranteed if the order is already in the process of being filled. A
// successful response indicates the cancellation request was accepted.
// See: https://docs.trading212.com/api/orders/cancelorder
func (op *Orders) CancelOrder(id int64) (*models.Empty, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.CancelOrder, id))
	return runOperation[models.Empty](op.api, http.MethodDelete, endpoint, nil).Object()
}

// GetPendingOrderByID operation.
// Retrieves a single pending order using its unique numerical ID.
// This is useful for checking the status of a specific order you have previously placed.
// See: https://docs.trading212.com/api/orders/orderbyid
func (op *Orders) GetPendingOrderByID(id int64) (*models.Order, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.GetPendingOrderByID, id))
	return runOperation[models.Order](op.api, http.MethodGet, endpoint, nil).Object()
}

// Positions operations
// Get a real-time overview of all your open positions, including quantity, average price, and current profit or loss.
// See: https://docs.trading212.com/api/positions
func (api *API) Positions() *Positions {
	return &Positions{api}
}

type Positions struct {
	api *API
}

// GetAllPositions operation.
// Fetch all open positions for your account.
// See: https://docs.trading212.com/api/positions/getpositions
func (op *Positions) GetAllPositions() (iter.Seq[models.Position], error) {
	return runOperation[models.Position](op.api, http.MethodGet, internal.GetAllPositions, nil).Items()
}

// HistoricalEvents operations
// Review your account's trading history. Access detailed records of past
// orders, dividend payments, and cash transactions, or generate downloadable
// CSV reports for analysis and record-keeping.
// See: https://docs.trading212.com/api/historical-events
func (api *API) HistoricalEvents() *HistoricalEvents {
	return &HistoricalEvents{api}
}

type HistoricalEvents struct {
	api *API
}

// GetPaidOutDividends operation.
// Get paid out dividends.
// See: https://docs.trading212.com/api/historical-events/dividends
func (op *HistoricalEvents) GetPaidOutDividends() (iter.Seq[models.Dividend], error) {
	return runOperation[models.Dividend](op.api, http.MethodGet, internal.GetDividends, nil).Items()
}

// GetHistoricalOrders operation.
// Get historical orders data.
// See: https://docs.trading212.com/api/historical-events/orders_1
func (op *HistoricalEvents) GetHistoricalOrders() (iter.Seq[models.OrderFill], error) {
	return runOperation[models.OrderFill](op.api, http.MethodGet, internal.GetHistoricalOrders, nil).Items()
}

// GetTransactions operation.
// Fetch superficial information about movements to and from your account.
// See: https://docs.trading212.com/api/historical-events/transactions
func (op *HistoricalEvents) GetTransactions() (iter.Seq[models.Transaction], error) {
	return runOperation[models.Transaction](op.api, http.MethodGet, internal.GetTransactions, nil).Items()
}

// ListReports operation.
// Retrieves a list of all requested CSV reports and their current status.
// Asynchronous Workflow:
//   - Call POST /history/exports to request a report. You will receive a reportId.
//   - Periodically call this endpoint (GET /history/exports) to check the status of the report corresponding to your reportId.
//   - Once the status is Finished, the downloadLink field will contain a URL to download the CSV file.
//
// See: https://docs.trading212.com/api/historical-events/getreports
func (op *HistoricalEvents) ListReports() (iter.Seq[models.Report], error) {
	return runOperation[models.Report](op.api, http.MethodGet, internal.ListReports, nil).Items()
}

// RequestReport operation.
// Initiates the generation of a CSV report containing historical account data.
// This is an asynchronous operation. The response will include a reportId which
// you can use to track the status of the generation process using the GET
// /history/exports endpoint.
// See: https://docs.trading212.com/api/historical-events/requestreport
func (op *HistoricalEvents) RequestReport(req models.ReportRequest) (*models.ReportId, error) {
	return runOperation[models.ReportId](op.api, http.MethodPost, internal.RequestReport, jsonBody{req}).Object()
}

// Pies operations
// Manage your investment Pies. Use these endpoints to create, view, update,
// and delete your custom portfolios, making automated and diversified investing simple.
//
// Deprecation notice: The current state of the Pies API, while still operational, won't be further supported and updated.
// See: https://docs.trading212.com/api/pies-(deprecated)
func (api *API) Pies() *Pies {
	return &Pies{api}
}

type Pies struct {
	api *API
}

// FetchAllPies operation.
// Fetches all pies for the account.
// See: https://docs.trading212.com/api/pies-(deprecated)/getall
func (op *Pies) FetchAllPies() (iter.Seq[models.PieSummary], error) {
	return runOperation[models.PieSummary](op.api, http.MethodGet, internal.GetAllPies, nil).Items()
}

// CreatePie operation.
// Creates a pie for the account by given params.
// See: https://docs.trading212.com/api/pies-(deprecated)/create
func (op *Pies) CreatePie(req models.PieRequest) (*models.PieDetails, error) {
	return runOperation[models.PieDetails](op.api, http.MethodPost, internal.CreatePie, req).Object()
}

// DeletePie operation.
// Deletes a pie by given id.
// See: https://docs.trading212.com/api/pies-(deprecated)/delete
func (op *Pies) DeletePie(id uint) error {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.DeletePie, id))
	return runOperation[models.Empty](op.api, http.MethodDelete, endpoint, nil).err
}

// FetchPie operation.
// Fetches a pies for the account with detailed information.
// See: https://docs.trading212.com/api/pies-(deprecated)/getdetailed
func (op *Pies) FetchPie(id uint) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.FetchPie, id))
	return runOperation[models.PieDetails](op.api, http.MethodGet, endpoint, nil).Object()
}

// UpdatePie operation.
// Updates a pie for the account by given params.
// See: https://docs.trading212.com/api/pies-(deprecated)/update
func (op *Pies) UpdatePie(id uint, req models.PieRequest) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.UpdatePie, id))
	return runOperation[models.PieDetails](op.api, http.MethodPost, endpoint, req).Object()
}

// DuplicatePies operation.
// Duplicates a pie for the account.
// See: https://docs.trading212.com/api/pies-(deprecated)/duplicatepie
func (op *Pies) DuplicatePies(id uint, req models.PieMetaRequest) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d/duplicate", internal.DuplicatePie, id))
	return runOperation[models.PieDetails](op.api, http.MethodPost, endpoint, req).Object()
}
