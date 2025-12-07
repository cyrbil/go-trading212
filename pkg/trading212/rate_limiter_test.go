package trading212

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestApplyRateLimit(t *testing.T) {
	t.Parallel()
	
	type args struct {
		path       string
		rateLimits map[string]APIRateLimits
	}
	tests := []struct {
		name  string
		args  args
		sleep bool
	}{
		{
			name: "ApplyRateLimit should do nothing on new path",
			args: args{
				path:       "new/path",
				rateLimits: make(map[string]APIRateLimits),
			},
			sleep: false,
		},
		{
			name: "ApplyRateLimit should do nothing on new path",
			args: args{
				path:       "new/path",
				rateLimits: make(map[string]APIRateLimits),
			},
			sleep: false,
		},
		{
			name: "ApplyRateLimit should do nothing when remaining left",
			args: args{
				path: "new/path",
				rateLimits: map[string]APIRateLimits{
					"new/path": {
						Remaining: 1,
					},
				},
			},
			sleep: false,
		},
		{
			name: "ApplyRateLimit should do nothing when reset happened",
			args: args{
				path: "new/path",
				rateLimits: map[string]APIRateLimits{
					"new/path": {
						Remaining: 0,
						Reset:     time.Now().Add(-5 * time.Minute),
					},
				},
			},
			sleep: false,
		},
		{
			name: "ApplyRateLimit should wait when no remaining and reset in future",
			args: args{
				path: "new/path",
				rateLimits: map[string]APIRateLimits{
					"new/path": {
						Remaining: 0,
						Reset:     time.Now().Add(5 * time.Minute),
					},
				},
			},
			sleep: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				rateLimiter := NewRateLimiter()
				rateLimiter.limits = tt.args.rateLimits

				called := false
				rateLimiter.sleep = func(_ time.Duration) {
					called = true
				}

				rateLimiter.ApplyRateLimit(tt.args.path)
				if called != tt.sleep {
					t.Errorf("ApplyRateLimit() called = %v; want %v", called, tt.sleep)
				}
			},
		)
	}
}

func TestParseRateLimits(t *testing.T) {
	t.Parallel()

	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    *APIRateLimits
		wantErr error
	}{
		{
			name:    "TestParseRateLimits should return error on nil response",
			args:    args{},
			want:    nil,
			wantErr: errHeaderNotFound,
		},
		{
			name: "TestParseRateLimits should return error on header not found",
			args: args{
				response: &http.Response{
					Header:  http.Header{},
					Request: &http.Request{URL: &url.URL{}},
				},
			},
			want:    nil,
			wantErr: errHeaderNotFound,
		},
		{
			name: "TestParseRateLimits should return error on non-int header",
			args: args{
				response: &http.Response{
					Header: http.Header{
						RateLimitHeaderLimit:     {"invalid"},
						RateLimitHeaderPeriod:    {"invalid"},
						RateLimitHeaderReset:     {"invalid"},
						RateLimitHeaderRemaining: {"invalid"},
						RateLimitHeaderUsed:      {"invalid"},
					},
					Request: &http.Request{URL: &url.URL{}},
				},
			},
			want:    nil,
			wantErr: errHeaderConversion,
		},
		{
			name: "TestParseRateLimits should return error on larger than int header period",
			args: args{
				response: &http.Response{
					Header: http.Header{
						RateLimitHeaderLimit:     {"0"},
						RateLimitHeaderPeriod:    {"9223372036854775808"},
						RateLimitHeaderReset:     {"0"},
						RateLimitHeaderRemaining: {"0"},
						RateLimitHeaderUsed:      {"0"},
					},
					Request: &http.Request{URL: &url.URL{}},
				},
			},
			want:    nil,
			wantErr: errHeaderConversion,
		},
		{
			name: "TestParseRateLimits should return error on larger than int header reset",
			args: args{
				response: &http.Response{
					Header: http.Header{
						RateLimitHeaderLimit:     {"0"},
						RateLimitHeaderPeriod:    {"0"},
						RateLimitHeaderReset:     {"9223372036854775808"},
						RateLimitHeaderRemaining: {"0"},
						RateLimitHeaderUsed:      {"0"},
					},
					Request: &http.Request{URL: &url.URL{}},
				},
			},
			want:    nil,
			wantErr: errHeaderConversion,
		},
		{
			name: "TestParseRateLimits with valid input",
			args: args{
				response: &http.Response{
					Header: http.Header{
						RateLimitHeaderLimit:     {"10"},
						RateLimitHeaderPeriod:    {"100"},
						RateLimitHeaderReset:     {"1765055699"},
						RateLimitHeaderRemaining: {"0"},
						RateLimitHeaderUsed:      {"0"},
					},
					Request: &http.Request{URL: &url.URL{}},
				},
			},
			want: &APIRateLimits{
				Limit:     10,
				Period:    100 * time.Second,
				Remaining: 0,
				Reset:     time.Unix(1765055699, 0),
				Used:      0,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				if tt.args.response != nil && tt.args.response.Header != nil {
					// so they are parsed properly
					for header, value := range tt.args.response.Header {
						tt.args.response.Header.Set(header, value[0])
					}
				}

				rateLimiter := NewRateLimiter()
				err := rateLimiter.ParseRateLimits("foo", tt.args.response)
				if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
					t.Errorf("ParseRateLimits() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				limits, ok := rateLimiter.limits["foo"]
				if (!ok && tt.want != nil) || (ok && !reflect.DeepEqual(limits, *tt.want)) {
					t.Errorf("ParseRateLimits() error limits not equal %v, want %v", limits, tt.want)
				}
			},
		)
	}
}
