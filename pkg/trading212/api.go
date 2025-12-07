package trading212

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 5 * time.Second

// APIURL string url for trading212 api.
type APIURL string

const (
	apiURLDemo APIURL = "https://demo.trading212.com"
	apiURLLive APIURL = "https://live.trading212.com" // gitleaks:allow # false positive
)

var (
	errEmptyDomain    = errors.New("domain should not be empty")
	errInvalidDomain  = errors.New("domain be a valid url")
	errEmptyAPIKey    = errors.New("API key should not be empty")
	errEmptyAPISecret = errors.New("API secret should not be empty")
)

// API client for trading212.
type API struct {
	*operations

	domain     *url.URL
	apiKey     string
	apiSecret  SecureString
	rateLimits *RateLimiter

	client *http.Client
}

// NewAPILive create a new client for trading212 API live.
func NewAPILive(apiKey string, apiSecret SecureString) (*API, error) {
	return NewAPI(apiURLLive, apiKey, apiSecret)
}

// NewAPIDemo create a new client for trading212 API demo.
func NewAPIDemo(apiKey string, apiSecret SecureString) (*API, error) {
	return NewAPI(apiURLDemo, apiKey, apiSecret)
}

// NewAPI create a new client for trading212 API.
func NewAPI(apiURL APIURL, apiKey string, apiSecret SecureString) (*API, error) {
	if apiURL == "" {
		return nil, errEmptyDomain
	}

	if apiKey == "" {
		return nil, errEmptyAPIKey
	}

	if apiSecret == "" {
		return nil, errEmptyAPISecret
	}

	domainURL, err := url.Parse(string(apiURL))
	if err != nil {
		return nil, errors.Join(errInvalidDomain, err)
	}

	api := &API{
		domain:     domainURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		rateLimits: NewRateLimiter(),
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       defaultTimeout,
		},
		operations: &operations{
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

	return api, nil
}
