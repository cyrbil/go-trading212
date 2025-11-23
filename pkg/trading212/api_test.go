package trading212

import (
	"reflect"
	"testing"
)

func TestNewAPI(t *testing.T) {
	type args struct {
		domain    APIDomain
		apiKey    string
		apiSecret SecureString
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "NewAPI should return an instance with attributes set from arguments",
			args: args{
				domain:    APIDomainDemo,
				apiKey:    "foo",
				apiSecret: "bar",
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := NewAPI(tt.args.domain, tt.args.apiKey, tt.args.apiSecret)
				if !reflect.DeepEqual(got.domain, tt.args.domain) {
					t.Errorf("NewAPI().domain = %v, want %v", got, tt.args.domain)
				}
				if !reflect.DeepEqual(got.apiKey, tt.args.apiKey) {
					t.Errorf("NewAPI().apiKey = %v, want %v", got, tt.args.apiKey)
				}
				if !reflect.DeepEqual(got.apiSecret, tt.args.apiSecret) {
					t.Errorf("NewAPI().apiSecret = %v, want %v", got, tt.args.apiSecret)
				}
				if got.rateLimits == nil {
					t.Errorf("NewAPI().rateLimits is nil")
				}
				if got.client.Timeout <= 0 {
					t.Errorf("NewAPI().client.Timeout = %v, want >0", got.client.Timeout)
				}
			},
		)
	}
}
