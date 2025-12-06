// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	internal "github.com/cyrbil/go-trading212/internal/pkg/trading212"
)

type mockRequestMaker struct {
	request IRequest
	err     error
}

func (mock *mockRequestMaker) NewRequest(_ string, _ internal.APIEndpoint, _ io.Reader) (IRequest, error) {
	return mock.request, mock.err
}

type mockIRequest struct {
	data        *json.RawMessage
	err         error
	httpRequest *http.Request
}

func (mock *mockIRequest) Do() (*json.RawMessage, error) {
	return mock.data, mock.err
}

func (mock *mockIRequest) http() *http.Request {
	return mock.httpRequest
}

func Test_runOperation(t *testing.T) {
	t.Run(
		"runOperation should return an error when the api fails to build request", func(t *testing.T) {
			api := &mockRequestMaker{err: errors.New("mocked error")}
			response := runOperation[any](api, "", "", nil)
			if response.err == nil {
				t.Error("expected error, got nil")
			}
		},
	)

	t.Run(
		"runOperation should return an error when the api fails to build request", func(t *testing.T) {
			api := &mockRequestMaker{request: &mockIRequest{err: errors.New("mocked error")}}
			response := runOperation[any](api, "", "", nil)
			if response.err == nil {
				t.Error("expected error, got nil")
			}
		},
	)
}

func newMockAPI(handler http.Handler) (*API, func(), error) {
	ts := httptest.NewTLSServer(handler)

	mockAPI, err := NewAPIDemo("foo", "bar")
	if err != nil {
		return nil, ts.Close, errors.Join(errors.New("error creating mock api"), err)
	}

	ts.Client().Timeout = mockAPI.client.Timeout
	mockAPI.client = ts.Client()
	mockAPI.domain, err = url.Parse(ts.URL)
	if err != nil {
		return nil, ts.Close, errors.Join(errors.New("error parsing mock url"), err)
	}

	return mockAPI, ts.Close, nil
}
func validateOperation[T any](t *testing.T, operation func(*API) (T, error), mockData string) T {
	mockAPI, terminate, err := newMockAPI(
		http.HandlerFunc(
			func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.Header().Set(internal.RateLimitHeaderLimit, "10")
				w.Header().Set(internal.RateLimitHeaderPeriod, "10")
				w.Header().Set(internal.RateLimitHeaderRemaining, "10")
				w.Header().Set(internal.RateLimitHeaderReset, "0")
				w.Header().Set(internal.RateLimitHeaderUsed, "0")
				w.WriteHeader(http.StatusOK)
				_, err := fmt.Fprintln(w, mockData)
				if err != nil {
					t.Fatalf("MockHTTP Write() error = %v", err)
				}
			},
		),
	)
	defer terminate()

	if err != nil {
		t.Fatalf("Error creating mock api; %v", err)
	}

	data, err := operation(mockAPI)
	if err != nil {
		t.Fatalf("Error calling operation; %v", err)
	}

	return data
}

func validateOperationObject[T any](t *testing.T, operation func(*API) (*T, error), mockData string) {
	data := validateOperation[*T](t, operation, mockData)
	if data == nil {
		t.Errorf("Error calling operation; data is nil")
	}
}

func validateOperationItems[T any](t *testing.T, operation func(*API) (iter.Seq[*T], error), mockData string) {
	data := validateOperation[iter.Seq[*T]](t, operation, mockData)
	if data == nil {
		t.Fatalf("Error calling operation; iterator is nil")
	}

	for item := range data {
		if item == nil {
			t.Errorf("Error calling operation; data[] contains a nil")
		}
	}
}
