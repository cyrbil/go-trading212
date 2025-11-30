package trading212

import (
	"net/http"
	"time"

	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

const defaultTimeout = 5 * time.Second

type APIDomain string

const (
	APIDomainDemo APIDomain = "demo.trading212.com"
	APIDomainLive APIDomain = "live.trading212.com"
)

type API struct {
	operations

	domain     APIDomain
	apiKey     string
	apiSecret  SecureString
	rateLimits map[string]internal.APIRateLimits

	client http.Client
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
		rateLimits: make(map[string]internal.APIRateLimits),
		client: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       defaultTimeout,
		},
		operations: operations{
			Account:          nil,
			Instruments:      nil,
			Orders:           nil,
			Positions:        nil,
			HistoricalEvents: nil,
			Pies:             nil,
		},
	}

	api.Account = &account{api}
	api.Instruments = &instruments{api}
	api.Orders = &orders{api}
	api.Positions = &positions{api}
	api.HistoricalEvents = &historicalEvents{api}
	api.Pies = &pies{api}

	return api
}
