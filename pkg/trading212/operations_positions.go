package trading212

import (
	"iter"
	"net/http"

	"github.com/cyrbil/go-trading212/pkg/trading212/models"
)

type positionsOperations interface {
	// GetAllPositions operation.
	// Fetch all open positions for your account.
	// See: https://docs.trading212.com/api/positions/getpositions
	GetAllPositions() (iter.Seq[*models.Position], error)
}

type positions struct {
	api requestMaker
}

func (op *positions) GetAllPositions() (iter.Seq[*models.Position], error) {
	return runOperation[models.Position](op.api, http.MethodGet, GetAllPositions, nil).Items()
}
