package trading212

import (
	"iter"
	"net/http"

	"github.com/cyrbil/go-trading212/pkg/trading212/models"
)

type operationGetExchangesMetadata interface {
	// GetExchangesMetadata operation.
	// Retrieves all accessible exchanges and their corresponding working schedules. Data is refreshed every 10 minutes.
	// See: https://docs.trading212.com/api/instruments/exchanges
	GetExchangesMetadata() (iter.Seq[*models.ExchangeMetadata], error)
}

type operationGetAllAvailableInstruments interface {
	// GetAllAvailableInstruments operation.
	// Retrieves all accessible instruments. Data is refreshed every 10 minutes.
	// See: https://docs.trading212.com/api/instruments/instruments
	GetAllAvailableInstruments() (iter.Seq[*models.Instrument], error)
}

type instrumentsOperations interface {
	operationGetExchangesMetadata
	operationGetAllAvailableInstruments
}

type instruments struct {
	api requestMaker
}

func (op *instruments) GetExchangesMetadata() (iter.Seq[*models.ExchangeMetadata], error) {
	return runOperation[models.ExchangeMetadata](op.api, http.MethodGet, GetExchangesMetadata, nil).Items()
}

func (op *instruments) GetAllAvailableInstruments() (iter.Seq[*models.Instrument], error) {
	return runOperation[models.Instrument](op.api, http.MethodGet, GetAllAvailableInstruments, nil).Items()
}
