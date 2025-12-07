package trading212

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	models "github.com/cyrbil/go-trading212/pkg/trading212/models"
	"reflect"
	"testing"
)

func Test_SecureString(t *testing.T) {
	t.Parallel()

	secureString := SecureString("foobar")

	//nolint:staticcheck
	secureStringFmt := fmt.Sprintf("%s", secureString)
	if secureStringFmt == "foobar" {
		t.Errorf("SecureString value is unprotected from format, got: %s", secureStringFmt)
	}

	secureStringJSON, err := json.Marshal(secureString)
	if err != nil {
		t.Error(err)
	}
	if string(secureStringJSON) == `"foobar"` {
		t.Errorf("SecureString value is unprotected from marshalling, got: %s", secureStringJSON)
	}

	if !reflect.DeepEqual(string(secureString), "foobar") {
		t.Errorf("SecureString value changed")
	}
}

func Test_jsonBody_Read(t *testing.T) {
	t.Parallel()

	type args struct {
		buf []byte
	}
	tests := []struct {
		name     string
		jsonBody *jsonBody
		args     args
		want     string
		wantErr  error
	}{
		{
			name: "jsonBody.Read() validations",
			jsonBody: newJSONBody(
				&models.PieMetaRequest{
					Icon: "foo",
					Name: "bar",
				},
			),
			args: args{
				buf: make([]byte, 30),
			},
			want:    `{"icon":"foo","name":"bar"}`,
			wantErr: nil,
		},
		{
			name: "jsonBody.Read() validations with json marshalling error",
			jsonBody: &jsonBody{
				data: &models.PieMetaRequest{
					Icon: "foo",
					Name: "bar",
				},
				marshaller: func(_ any) ([]byte, error) {
					return nil, errors.New("mock error")
				},
				reader: nil,
			},
			args: args{
				buf: make([]byte, 30),
			},
			want:    `{"icon":"foo","name":"bar"}`,
			wantErr: errConversionBody,
		},
		{
			name: "jsonBody.Read() validations with json reading error",
			jsonBody: &jsonBody{
				data: &models.PieMetaRequest{
					Icon: "foo",
					Name: "bar",
				},
				marshaller: json.Marshal,
				reader: func(_ []byte, _ []byte) (int, error) {
					return 0, errors.New("mock error")
				},
			},
			args: args{
				buf: make([]byte, 30),
			},
			want:    `{"icon":"foo","name":"bar"}`,
			wantErr: errConversionBody,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := tt.jsonBody.Read(tt.args.buf)
				if err != nil {
					if !errors.Is(err, tt.wantErr) {
						t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}

				if got < len(tt.want) {
					t.Errorf("jsonBody.Read() length error; got = %v, want %v", got, len(tt.want))
				}

				if !bytes.HasPrefix(tt.args.buf, []byte(tt.want)) {
					t.Errorf("jsonBody.Read() data error; got = %s, want %v", tt.args.buf, tt.want)
				}
			},
		)
	}
}
