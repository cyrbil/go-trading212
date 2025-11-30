package trading212

import (
	"net/http"
	"time"
)

import (
	"github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type APIDomain string

//goland:noinspection GoUnusedConst
const (
	APIDomainDemo APIDomain = "demo.trading212.com"
	APIDomainLive APIDomain = "live.trading212.com"
)

type API struct {
	domain     APIDomain
	apiKey     string
	apiSecret  SecureString
	rateLimits map[string]trading212.APIRateLimits

	client http.Client

	operations
}

func NewAPILive(apiKey string, apiSecret SecureString) *API {
	return NewAPI(APIDomainLive, apiKey, apiSecret)
}

func NewAPIDemo(apiKey string, apiSecret SecureString) *API {
	return NewAPI(APIDomainDemo, apiKey, apiSecret)
}

func NewAPI(domain APIDomain, apiKey string, apiSecret SecureString) *API {
	api := &API{
		domain:     domain,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		rateLimits: make(map[string]trading212.APIRateLimits),
		client:     http.Client{Timeout: 5 * time.Second},
	}

	api.Account = &account{api}
	api.Instruments = &instruments{api}
	api.Orders = &orders{api}
	api.Positions = &positions{api}
	api.HistoricalEvents = &historicalEvents{api}
	api.Pies = &pies{api}

	return api
}
