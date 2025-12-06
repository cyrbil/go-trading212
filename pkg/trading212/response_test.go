package trading212

import (
	"encoding/json"
	"errors"
	"testing"
)

func Test_Response_validate(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name    string
		r       Response[T]
		wantErr bool
	}
	tests := []testCase[any]{
		{
			name: "Response_validate should return inner error",
			r: Response[any]{
				err:     errors.New("mock error"),
				request: nil,
				raw:     nil,
			},
			wantErr: true,
		},
		{
			name: "Response_validate should return error if no request",
			r: Response[any]{
				err:     nil,
				request: nil,
				raw:     &json.RawMessage{},
			},
			wantErr: true,
		},
		{
			name: "Response_validate should return error if no data",
			r: Response[any]{
				err:     nil,
				request: &Request{},
				raw:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				if err := tt.r.validate(); (err != nil) != tt.wantErr {
					t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
