// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"errors"
	"testing"
)

func validateAPI(prefix string, t *testing.T, api *API, err error) {
	// check api
	if api == nil {
		t.Errorf("%s, api is nil", prefix)
		return
	}
	if err != nil {
		t.Errorf("%s, error returned, err = %v", prefix, err)
	}

	// check fields
	if api.apiKey == "" {
		t.Errorf("%s, apiKey is empty", prefix)
	}
	if api.apiSecret == "" {
		t.Errorf("%s, apiSecret is empty", prefix)
	}

	// check data structures
	if api.domain == nil {
		t.Errorf("%s, domain is nil", prefix)
	}
	if api.rateLimits == nil {
		t.Errorf("%s, rateLimits is nil", prefix)
	}
	if api.client == nil {
		t.Errorf("%s, http client is nil", prefix)
	}

	// check embedded operations
	if api.operations == nil {
		t.Errorf("%s, operations is nil", prefix)
	}
	if api.Account == nil {
		t.Errorf("%s, Account is nil", prefix)
	}
	if api.Instruments == nil {
		t.Errorf("%s, Instruments is nil", prefix)
	}
	if api.Orders == nil {
		t.Errorf("%s, Orders is nil", prefix)
	}
	if api.Positions == nil {
		t.Errorf("%s, Positions is nil", prefix)
	}
	if api.HistoricalEvents == nil {
		t.Errorf("%s, HistoricalEvents is nil", prefix)
	}
	if api.Pies == nil {
		t.Errorf("%s, Pies is nil", prefix)
	}
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
				got, err := NewAPI(tt.args.domain, tt.args.apiKey, tt.args.apiSecret)
				if tt.err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPI() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}
					return
				}

				validateAPI("NewAPI() api validation error", t, got, nil)
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
				got, err := NewAPIDemo(tt.args.apiKey, tt.args.apiSecret)
				if err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPIDemo() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}
					return
				}

				validateAPI("NewAPIDemo() api validation error", t, got, nil)
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
				got, err := NewAPILive(tt.args.apiKey, tt.args.apiSecret)
				if err != nil {
					if !errors.Is(err, tt.err) {
						t.Errorf("NewAPILive() returned an unexpected error; expect: %v, got: %v", tt.err, err)
					}
					return
				}

				validateAPI("NewAPILive() api validation error", t, got, nil)
				if got.domain.String() != string(apiURLLive) {
					t.Errorf("NewAPILive() should return the Live domain")
				}
			},
		)
	}
}
