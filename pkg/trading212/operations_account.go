package trading212

import (
	"net/http"

	"github.com/cyrbil/go-trading212/pkg/trading212/models"
)

type accountOperations interface {
	// GetAccountSummary operation.
	// Provides a breakdown of your account's cash and investment metrics,
	// including available funds, invested capital, and total account value.
	// See: https://docs.trading212.com/api/accounts/getaccountsummary
	GetAccountSummary() (*models.AccountSummary, error)
}

type account struct {
	api requestMaker
}

func (op *account) GetAccountSummary() (*models.AccountSummary, error) {
	return runOperation[models.AccountSummary](op.api, http.MethodGet, GetAccountSummary, nil).Object()
}
