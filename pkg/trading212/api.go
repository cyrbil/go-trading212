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
	APIDomainLive APIDomain = "live.trading212.com"  // gitleaks:allow # false positive
)

type API struct {
	domain     APIDomain
	apiKey     string
	apiSecret  SecureString
	rateLimits map[string]trading212.APIRateLimits

	client http.Client
}

func NewAPI(domain APIDomain, apiKey string, apiSecret SecureString) *API {
	return &API{
		domain:     domain,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		rateLimits: make(map[string]trading212.APIRateLimits),
		client:     http.Client{Timeout: 5 * time.Second},
	}
}
