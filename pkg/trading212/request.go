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

var (
	errRequestTimeout = errors.New("request timeout")
	errNewHttp        = fmt.Errorf("fail to create http request")
	errApiRequest     = errors.New("error executing api request")
	errReadingApi     = errors.New("error reading api response")
)

type Request struct {
	//nolint:containedctx
	Ctx         context.Context
	cancel      context.CancelCauseFunc
	api         *API
	httpRequest *http.Request
	retries     int
}

type requestMaker interface {
	NewRequest(method string, path internal.APIEndpoint, body io.Reader) (*Request, error)
}

// NewRequest build a Request for the API.
// Prefer to use the available methods instead.
func (api *API) NewRequest(method string, path internal.APIEndpoint, body io.Reader) (*Request, error) {
	endpoint := fmt.Sprintf("https://%s/%s", api.domain, path)

	ctx := context.Background()
	ctx, timeoutCancel := context.WithTimeoutCause(ctx, api.client.Timeout, errRequestTimeout)
	ctx, causeCancel := context.WithCancelCause(ctx)
	cancel := func(cause error) {
		timeoutCancel()
		causeCancel(cause)
	}

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

	return &Request{Ctx: ctx, cancel: cancel, api: api, httpRequest: request, retries: 0}, nil
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
		request.cancel(err)

		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			request.cancel(err)
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

	if response.StatusCode == internal.RateLimitedErrorCode {
		time.Sleep(time.Duration(request.retries) * time.Second)

		request.retries++

		return request.Do()
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		err := fmt.Errorf("error api return non 200, status: %s", response.Status)
		request.cancel(err)

		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		err := errors.Join(errReadingApi, err)
		request.cancel(err)

		return nil, err
	}

	slog.Debug("Response body", "body", data)

	return (*json.RawMessage)(&data), nil
}
