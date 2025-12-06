package trading212

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const defaultMaxRetries = 10

var (
	errNewHTTP    = errors.New("fail to create http request")
	errAPIRequest = errors.New("error executing api request")
	errReadingAPI = errors.New("error reading api response")
	errNon200     = errors.New("error api return non http 200")
	errHTTP401    = errors.New("error api return http 401; Bad API key")
	errHTTP403    = errors.New("error api return http 403; Scope missing for API key")
	errHTTP408    = errors.New("error api return http 408; Timed-out")
	errHTTP429    = errors.New("error api return http 429; Rate-Limited")
)

type knownErrorCode int

const (
	badAPIKey    knownErrorCode = 401
	scopeMissing knownErrorCode = 403
	timeout      knownErrorCode = 408
	rateLimited  knownErrorCode = 429
)

func httpError(code int, status string) error {
	var err error

	switch knownErrorCode(code) {
	case badAPIKey:
		err = errHTTP401
	case scopeMissing:
		err = errHTTP403
	case timeout:
		err = errHTTP408
	case rateLimited:
		err = errHTTP429
	default:
		err = errNon200
	}

	return fmt.Errorf("%w (status: %s)", err, status)
}

// IRequest Request interface.
type IRequest interface {
	Do() (*json.RawMessage, error)
	http() *http.Request
}

// Request API request.
type Request struct {
	//nolint:containedctx
	Ctx         context.Context
	cancel      context.CancelCauseFunc
	api         *API
	httpRequest *http.Request
	retries     int
	maxRetries  int
}

type requestMaker interface {
	NewRequest(method string, path APIEndpoint, body io.Reader) (IRequest, error)
}

// NewRequest build a Request for the API.
// Prefer to use the available methods instead.
//
//nolint:ireturn
func (api *API) NewRequest(method string, path APIEndpoint, body io.Reader) (IRequest, error) {
	endpoint := api.domain.JoinPath(string(path)).String()

	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)

	request, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		cancel(err)

		return nil, errors.Join(errNewHTTP, err)
	}

	// authentication
	request.SetBasicAuth(api.apiKey, string(api.apiSecret))
	// api accepts json
	request.Header.Set("Content-Type", "application/json")
	// extend default pagination from 20 to 50 when available
	query := request.URL.Query()
	query.Set("limit", "50")
	request.URL.RawQuery = query.Encode()

	return &Request{
		Ctx:         ctx,
		cancel:      cancel,
		api:         api,
		httpRequest: request,
		retries:     0,
		maxRetries:  defaultMaxRetries,
	}, nil
}

// Do executes the current request.
func (request *Request) Do() (*json.RawMessage, error) {
	defer request.cancel(nil)

	rateLimitPath := request.httpRequest.URL.EscapedPath()
	request.api.rateLimits.ApplyRateLimit(rateLimitPath)

	//nolint:bodyclose // body is closed in lambda
	response, err := request.api.client.Do(request.httpRequest)
	if err != nil {
		err := errors.Join(errAPIRequest, err)

		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Warn("error closing api response body")
		}
	}(response.Body)

	slog.Debug("Request status", "status", response.Status)

	err = request.api.rateLimits.ParseRateLimits(rateLimitPath, response)
	if err != nil {
		slog.Warn("Fail to parse rate limits", "error", err)
	}

	if response.StatusCode == int(rateLimited) || response.StatusCode == int(timeout) {
		request.retries++
		if request.retries < request.maxRetries {
			time.Sleep(time.Duration(request.retries) * time.Second)

			return request.Do()
		}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		err := httpError(response.StatusCode, response.Status)

		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		err := errors.Join(errReadingAPI, err)

		return nil, err
	}

	slog.Debug("Response body", "body", data)

	return (*json.RawMessage)(&data), nil
}

func (request *Request) http() *http.Request {
	return request.httpRequest
}
