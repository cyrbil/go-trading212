// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"errors"
	"fmt"
	"testing"
)

func validateAPI(api *API, err error) error {
	// check api
	if api == nil {
		return errors.New("api is nil")
	}

	if err != nil {
		return fmt.Errorf("error returned, err = %w", err)
	}

	// check fields
	if api.apiKey == "" {
		return errors.New("apiKey is empty")
	}

	if api.apiSecret == "" {
		return errors.New("apiSecret is empty")
	}

	// check data structures
	if api.domain == nil {
		return errors.New("domain is nil")
	}

	if api.rateLimits == nil {
		return errors.New("rateLimits is nil")
	}

	if api.client == nil {
		return errors.New("http client is nil")
	}

	// check embedded operations
	if api.operations == nil {
		return errors.New("operations is nil")
	}

	if api.Account == nil {
		return errors.New("api.Account is nil")
	}

	if api.Instruments == nil {
		return errors.New("api.Instruments is nil")
	}

	if api.Orders == nil {
		return errors.New("api.Orders is nil")
	}

	if api.Positions == nil {
		return errors.New("api.Positions is nil")
	}

	if api.HistoricalEvents == nil {
		return errors.New("api.HistoricalEvents is nil")
	}

	if api.Pies == nil {
		return errors.New("api.Pies is nil")
	}

	return nil
}

func Test_NewAPI(t *testing.T) {
	t.Parallel()

	type args struct {
		domain    APIURL
		apiKey    string
		apiSecret SecureString
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Validate NewAPI()",
			args: args{
				domain:    "foo.bar",
				apiKey:    "foo",
				apiSecret: "bar",
			},
			err: nil,
		},
		{
			name: "Validate NewAPI() with empty domain",
			args: args{
				domain:    "",
				apiKey:    "foo",
				apiSecret: "bar",
			},
			err: errEmptyDomain,
		},
		{
			name: "Validate NewAPI() with invalid domain",
			args: args{
				domain:    "invalid%x99",
				apiKey:    "foo",
				apiSecret: "bar",
			},
			err: errInvalidDomain,
		},
		{
			name: "Validate NewAPI() with invalid key",
			args: args{
				domain:    "foo.bar",
				apiKey:    "",
				apiSecret: "bar",
			},
			err: errEmptyAPIKey,
		},
		{
			name: "Validate NewAPI() with invalid secret",
			args: args{
				domain:    "foo.bar",
				apiKey:    "foo",
				apiSecret: "",
			},
			err: errEmptyAPISecret,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				got, err := NewAPI(tt.args.domain, tt.args.apiKey, tt.args.apiSecret)
				if tt.err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPI() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}
					return
				}

				err = validateAPI(got, err)
				if err != nil {
					t.Errorf("NewAPI() returned an unexpected error; %v", err)
					return
				}

				if got.domain.String() != string(tt.args.domain) {
					t.Errorf(
						"NewAPIDemo() should return the given domain; expect: %s, got: %s", tt.args.domain, got.domain,
					)
				}
			},
		)
	}
}

func Test_NewAPIDemo(t *testing.T) {
	t.Parallel()

	type args struct {
		apiKey    string
		apiSecret SecureString
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Validate NewAPIDemo()",
			args: args{
				apiKey:    "foo",
				apiSecret: "bar",
			},
			err: nil,
		},
		{
			name: "Validate NewAPIDemo() with invalid key",
			args: args{
				apiKey:    "",
				apiSecret: "bar",
			},
			err: errEmptyAPIKey,
		},
		{
			name: "Validate NewAPIDemo() with invalid secret",
			args: args{
				apiKey:    "foo",
				apiSecret: "",
			},
			err: errEmptyAPISecret,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				got, err := NewAPIDemo(tt.args.apiKey, tt.args.apiSecret)
				if err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPIDemo() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}

					return
				}

				err = validateAPI(got, err)
				if err != nil {
					t.Errorf("NewAPIDemo() returned an unexpected error; %v", err)

					return
				}

				if got.domain.String() != string(apiURLDemo) {
					t.Errorf("NewAPIDemo() should return the Demo domain")
				}
			},
		)
	}
}

func Test_NewAPILive(t *testing.T) {
	t.Parallel()

	type args struct {
		apiKey    string
		apiSecret SecureString
	}

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Validate NewAPILive()",
			args: args{
				apiKey:    "foo",
				apiSecret: "bar",
			},
			err: nil,
		},
		{
			name: "Validate NewAPILive() with invalid key",
			args: args{
				apiKey:    "",
				apiSecret: "bar",
			},
			err: errEmptyAPIKey,
		},
		{
			name: "Validate NewAPILive() with invalid secret",
			args: args{
				apiKey:    "foo",
				apiSecret: "",
			},
			err: errEmptyAPISecret,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()
				got, err := NewAPILive(tt.args.apiKey, tt.args.apiSecret)
				if err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPILive() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}
					return
				}

				err = validateAPI(got, err)
				if err != nil {
					t.Errorf("NewAPILive() returned an unexpected error; %v", err)
					return
				}

				if got.domain.String() != string(apiURLLive) {
					t.Errorf("NewAPILive() should return the Live domain")
				}
			},
		)
	}
}
