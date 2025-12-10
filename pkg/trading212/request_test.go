package trading212

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestAPI_NewRequest(t *testing.T) {
	type fields struct {
		operations *operations
		domain     *url.URL
		apiKey     string
		apiSecret  SecureString
		rateLimits *RateLimiter
		client     *http.Client
	}
	type args struct {
		method string
		path   APIEndpoint
		body   io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    IRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				operations: tt.fields.operations,
				domain:     tt.fields.domain,
				apiKey:     tt.fields.apiKey,
				apiSecret:  tt.fields.apiSecret,
				rateLimits: tt.fields.rateLimits,
				client:     tt.fields.client,
			}
			got, err := api.NewRequest(tt.args.method, tt.args.path, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Do(t *testing.T) {
	type fields struct {
		Ctx         context.Context
		cancel      context.CancelCauseFunc
		api         *API
		httpRequest *http.Request
		retries     int
		maxRetries  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *json.RawMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &Request{
				Ctx:         tt.fields.Ctx,
				cancel:      tt.fields.cancel,
				api:         tt.fields.api,
				httpRequest: tt.fields.httpRequest,
				retries:     tt.fields.retries,
				maxRetries:  tt.fields.maxRetries,
			}
			got, err := request.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Request_httpError(t *testing.T) {
	t.Parallel()

	type args struct {
		code   int
		status string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "httpError should return wrapped 401",
			args: args{
				code:   401,
				status: http.StatusText(401),
			},
			want: errHTTP401,
		},
		{
			name: "httpError should return wrapped 403",
			args: args{
				code:   403,
				status: http.StatusText(403),
			},
			want: errHTTP403,
		},
		{
			name: "httpError should return wrapped 408",
			args: args{
				code:   408,
				status: http.StatusText(408),
			},
			want: errHTTP408,
		},
		{
			name: "httpError should return wrapped 429",
			args: args{
				code:   429,
				status: http.StatusText(429),
			},
			want: errHTTP429,
		},
		{
			name: "httpError should return wrapped 404",
			args: args{
				code:   404,
				status: http.StatusText(404),
			},
			want: errNon200,
		},
		{
			name: "httpError should return wrapped 500",
			args: args{
				code:   505,
				status: http.StatusText(500),
			},
			want: errNon200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := httpError(tt.args.code, tt.args.status)
			if !errors.Is(err, tt.want) {
				t.Errorf("httpError() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func Test_httpError(t *testing.T) {
	type args struct {
		code   int
		status string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := httpError(tt.args.code, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("httpError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
