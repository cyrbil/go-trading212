package trading212

import (
	"fmt"
	"iter"
	"net/http"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type operationFetchAllPies interface {
	// FetchAllPies operation.
	// Fetches all pies for the account.
	// See: https://docs.trading212.com/api/pies-(deprecated)/getall
	FetchAllPies() (iter.Seq[models.PieSummary], error)
}

type operationCreatePie interface {
	// CreatePie operation.
	// Creates a pie for the account by given params.
	// See: https://docs.trading212.com/api/pies-(deprecated)/create
	CreatePie(req models.PieRequest) (*models.PieDetails, error)
}

type operationDeletePie interface {
	// DeletePie operation.
	// Deletes a pie by given id.
	// See: https://docs.trading212.com/api/pies-(deprecated)/delete
	DeletePie(id uint) error
}

type operationFetchPie interface {
	// FetchPie operation.
	// Fetches a pies for the account with detailed information.
	// See: https://docs.trading212.com/api/pies-(deprecated)/getdetailed
	FetchPie(id uint) (*models.PieDetails, error)
}

type operationUpdatePie interface {
	// UpdatePie operation.
	// Updates a pie for the account by given params.
	// See: https://docs.trading212.com/api/pies-(deprecated)/update
	UpdatePie(id uint, req models.PieRequest) (*models.PieDetails, error)
}

type operationDuplicatePies interface {
	// DuplicatePies operation.
	// Duplicates a pie for the account.
	// See: https://docs.trading212.com/api/pies-(deprecated)/duplicatepie
	DuplicatePies(id uint, req models.PieMetaRequest) (*models.PieDetails, error)
}

type piesOperations interface {
	operationFetchAllPies
	operationCreatePie
	operationDeletePie
	operationFetchPie
	operationUpdatePie
	operationDuplicatePies
}

type pies struct {
	api requestMaker
}

func (op *pies) FetchAllPies() (iter.Seq[models.PieSummary], error) {
	return runOperation[models.PieSummary](op.api, http.MethodGet, internal.GetAllPies, nil).Items()
}

func (op *pies) CreatePie(req models.PieRequest) (*models.PieDetails, error) {
	return runOperation[models.PieDetails](op.api, http.MethodPost, internal.CreatePie, req).Object()
}

func (op *pies) DeletePie(id uint) error {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.DeletePie, id))

	return runOperation[models.Empty](op.api, http.MethodDelete, endpoint, nil).err
}

func (op *pies) FetchPie(id uint) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.FetchPie, id))

	return runOperation[models.PieDetails](op.api, http.MethodGet, endpoint, nil).Object()
}

func (op *pies) UpdatePie(id uint, req models.PieRequest) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d", internal.UpdatePie, id))

	return runOperation[models.PieDetails](op.api, http.MethodPost, endpoint, req).Object()
}

func (op *pies) DuplicatePies(id uint, req models.PieMetaRequest) (*models.PieDetails, error) {
	endpoint := internal.APIEndpoint(fmt.Sprintf("%s/%d/duplicate", internal.DuplicatePie, id))

	return runOperation[models.PieDetails](op.api, http.MethodPost, endpoint, req).Object()
}
