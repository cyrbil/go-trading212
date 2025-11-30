package trading212

import (
	"iter"
	"net/http"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type positionsOperations interface {
	// GetAllPositions operation.
	// Fetch all open positions for your account.
	// See: https://docs.trading212.com/api/positions/getpositions
	GetAllPositions() (iter.Seq[models.Position], error)
}

type positions struct {
	api requestMaker
}

func (op *positions) GetAllPositions() (iter.Seq[models.Position], error) {
	return runOperation[models.Position](op.api, http.MethodGet, internal.GetAllPositions, nil).Items()
}
