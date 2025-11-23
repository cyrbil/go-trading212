package trading212

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

import (
	"github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type Request struct {
	api         *API
	httpRequest *http.Request
	retries     int
}

// NewRequest build a Request for the API.
// Prefer to use the available methods instead.
func (api *API) NewRequest(method string, path trading212.APIEndpoint, body io.Reader) (*Request, error) {
	endpoint := fmt.Sprintf("https://%s/%s", api.domain, path)

	request, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}
	// authentication
	request.SetBasicAuth(api.apiKey, string(api.apiSecret))
	// api accepts json
	request.Header.Set("Content-Type", "application/json")
	// extend default pagination from 20 to 50 when available
	request.URL.Query().Set("limit", "50")

	return &Request{api, request, 0}, nil
}

// Do executes the current request.
func (request *Request) Do() (*json.RawMessage, error) {
	rateLimitPath := request.httpRequest.URL.EscapedPath()
	trading212.ApplyRateLimit(rateLimitPath, request.api.rateLimits)

	response, err := request.api.client.Do(request.httpRequest)
	if err != nil {
		err := errors.Join(errors.New("error executing api request"), err)
		return nil, err
	}
	slog.Debug("Request status", "status", response.Status)
	limits, err := trading212.ParseRateLimits(response)
	if err != nil {
		return nil, err
	}
	request.api.rateLimits[rateLimitPath] = *limits

	if response.StatusCode == 429 {
		time.Sleep(time.Duration(request.retries) * time.Second)
		request.retries += 1
		return request.Do()
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		err := fmt.Errorf("error api return non 200, status: %s", response.Status)
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		err := errors.Join(errors.New("error reading api response"), err)
		return nil, err
	}
	slog.Debug("Response body", "body", data)

	return (*json.RawMessage)(&data), nil
}
