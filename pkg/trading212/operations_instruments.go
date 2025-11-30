package trading212

import (
	"iter"
	"net/http"
)

import (
	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type operationGetExchangesMetadata interface {
	// GetExchangesMetadata operation.
	// Retrieves all accessible exchanges and their corresponding working schedules. Data is refreshed every 10 minutes.
	// See: https://docs.trading212.com/api/instruments/exchanges
	GetExchangesMetadata() (iter.Seq[models.ExchangeMetadata], error)
}

type operationGetAllAvailableInstruments interface {
	// GetAllAvailableInstruments operation.
	// Retrieves all accessible instruments. Data is refreshed every 10 minutes.
	// See: https://docs.trading212.com/api/instruments/instruments
	GetAllAvailableInstruments() (iter.Seq[models.Instrument], error)
}

type instrumentsOperations interface {
	operationGetExchangesMetadata
	operationGetAllAvailableInstruments
}

type instruments struct {
	api requestMaker
}

func (op *instruments) GetExchangesMetadata() (iter.Seq[models.ExchangeMetadata], error) {
	return runOperation[models.ExchangeMetadata](op.api, http.MethodGet, internal.GetExchangesMetadata, nil).Items()
}

func (op *instruments) GetAllAvailableInstruments() (iter.Seq[models.Instrument], error) {
	return runOperation[models.Instrument](op.api, http.MethodGet, internal.GetAllAvailableInstruments, nil).Items()
}
