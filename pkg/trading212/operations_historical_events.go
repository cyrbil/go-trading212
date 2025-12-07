package trading212

import (
	"iter"
	"net/http"

	"github.com/cyrbil/go-trading212/pkg/trading212/models"
)

type operationGetPaidOutDividends interface {
	// GetPaidOutDividends operation.
	// Get paid out dividends.
	// See: https://docs.trading212.com/api/historical-events/dividends
	GetPaidOutDividends() (iter.Seq[*models.Dividend], error)
}

type operationGetHistoricalOrders interface {
	// GetHistoricalOrders operation.
	// Get historical orders data.
	// See: https://docs.trading212.com/api/historical-events/orders_1
	GetHistoricalOrders() (iter.Seq[*models.OrderFill], error)
}

type operationGetTransactions interface {
	// GetTransactions operation.
	// Fetch superficial information about movements to and from your account.
	// See: https://docs.trading212.com/api/historical-events/transactions
	GetTransactions() (iter.Seq[*models.Transaction], error)
}

type operationListReports interface {
	// ListReports operation.
	// Retrieves a list of all requested CSV reports and their current status.
	// Asynchronous Workflow:
	//   - Call POST /history/exports to request a report. You will receive a reportId.
	//   - Periodically call this endpoint (GET /history/exports) to check the status
	//     of the report corresponding to your reportId.
	//   - Once the status is Finished, the downloadLink field will contain a URL to download the CSV file.
	//
	// See: https://docs.trading212.com/api/historical-events/getreports
	ListReports() (iter.Seq[*models.Report], error)
}

type operationRequestReport interface {
	// RequestReport operation.
	// Initiates the generation of a CSV report containing historical account data.
	// This is an asynchronous operation. The response will include a reportId which
	// you can use to track the status of the generation process using the GET
	// /history/exports endpoint.
	// See: https://docs.trading212.com/api/historical-events/requestreport
	RequestReport(req models.ReportRequest) (*models.ReportID, error)
}

type historicalEventsOperations interface {
	operationGetPaidOutDividends
	operationGetHistoricalOrders
	operationGetTransactions
	operationListReports
	operationRequestReport
}

type historicalEvents struct {
	api requestMaker
}

func (op *historicalEvents) GetPaidOutDividends() (iter.Seq[*models.Dividend], error) {
	return runOperation[models.Dividend](op.api, http.MethodGet, GetDividends, nil).Items()
}

func (op *historicalEvents) GetHistoricalOrders() (iter.Seq[*models.OrderFill], error) {
	return runOperation[models.OrderFill](op.api, http.MethodGet, GetHistoricalOrders, nil).Items()
}

func (op *historicalEvents) GetTransactions() (iter.Seq[*models.Transaction], error) {
	return runOperation[models.Transaction](op.api, http.MethodGet, GetTransactions, nil).Items()
}

func (op *historicalEvents) ListReports() (iter.Seq[*models.Report], error) {
	return runOperation[models.Report](op.api, http.MethodGet, ListReports, nil).Items()
}

func (op *historicalEvents) RequestReport(req models.ReportRequest) (*models.ReportID, error) {
	return runOperation[models.ReportID](op.api, http.MethodPost, RequestReport, req).Object()
}
