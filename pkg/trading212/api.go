package trading212

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

const defaultTimeout = 5 * time.Second

type APIURL string

const (
	APIURLDemo APIURL = "https://demo.trading212.com"
	APIURLLive APIURL = "https://live.trading212.com"
)

var (
	errEmptyDomain    = errors.New("domain should not be empty")
	errInvalidDomain  = errors.New("domain be a valid url")
	errEmptyApiKey    = errors.New("API key should not be empty")
	errEmptyApiSecret = errors.New("API secret should not be empty")
)

type API struct {
	*operations

	domain     *url.URL
	apiKey     string
	apiSecret  SecureString
	rateLimits map[string]internal.APIRateLimits

	client *http.Client
}

func NewAPILive(apiKey string, apiSecret SecureString) (*API, error) {
	return NewAPI(APIURLLive, apiKey, apiSecret)
}

func NewAPIDemo(apiKey string, apiSecret SecureString) (*API, error) {
	return NewAPI(APIURLDemo, apiKey, apiSecret)
}

func NewAPI(apiUrl APIURL, apiKey string, apiSecret SecureString) (*API, error) {
	if apiUrl == "" {
		return nil, errEmptyDomain
	}
	if apiKey == "" {
		return nil, errEmptyApiKey
	}
	if apiSecret == "" {
		return nil, errEmptyApiSecret
	}

	domainUrl, err := url.Parse(string(apiUrl))
	if err != nil {
		return nil, errors.Join(errInvalidDomain, err)
	}

	api := &API{
		domain:     domainUrl,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		rateLimits: make(map[string]internal.APIRateLimits),
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
