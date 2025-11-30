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

	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

const defaultMaxRetries = 10

var (
	errNewHttp    = errors.New("fail to create http request")
	errApiRequest = errors.New("error executing api request")
	errReadingApi = errors.New("error reading api response")
	errNon200     = errors.New("error api return non http 200")
	errHttp401    = errors.New("error api return http 401; Bad API key")
	errHttp403    = errors.New("error api return http 403; Scope missing for API key")
	errHttp408    = errors.New("error api return http 408; Timed-out")
)

type knownErrorCode int

const (
	badApiKey    knownErrorCode = 401
	scopeMissing knownErrorCode = 403
	timeout      knownErrorCode = 408
	rateLimited  knownErrorCode = 429
)

func httpError(code int, status string) error {
	var err error

	switch knownErrorCode(code) {
	case badApiKey:
		err = errHttp401
		break
	case scopeMissing:
		err = errHttp403
		break
	case timeout:
		err = errHttp408
		break
	default:
		err = errNon200
		break
	}

	return fmt.Errorf("%w (status: %s)", err, status)
}

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
	NewRequest(method string, path internal.APIEndpoint, body io.Reader) (IRequest, error)
}

// NewRequest build a Request for the API.
// Prefer to use the available methods instead.
func (api *API) NewRequest(method string, path internal.APIEndpoint, body io.Reader) (IRequest, error) {
	endpoint := api.domain.JoinPath(string(path)).String()

	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)

	request, err := http.NewRequestWithContext(ctx, method, endpoint, body)
	if err != nil {
		cancel(err)

		return nil, errors.Join(errNewHttp, err)
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

type IRequest interface {
	Do() (*json.RawMessage, error)
	http() *http.Request
}

// Do executes the current request.
func (request *Request) Do() (*json.RawMessage, error) {
	defer request.cancel(nil)

	rateLimitPath := request.httpRequest.URL.EscapedPath()
	internal.ApplyRateLimit(rateLimitPath, request.api.rateLimits)

	//nolint:bodyclose // body is closed in lambda
	response, err := request.api.client.Do(request.httpRequest)
	if err != nil {
		err := errors.Join(errApiRequest, err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Warn("error closing api response body")
		}
	}(response.Body)

	slog.Debug("Request status", "status", response.Status)

	limits, err := internal.ParseRateLimits(response)
	if err != nil {
		slog.Warn("Fail to parse rate limits", "error", err)
	} else {
		request.api.rateLimits[rateLimitPath] = *limits
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
		err := errors.Join(errReadingApi, err)
		return nil, err
	}

	slog.Debug("Response body", "body", data)

	return (*json.RawMessage)(&data), nil
}

func (request *Request) http() *http.Request {
	return request.httpRequest
}
