package trading212

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// SecureString is a string type that doesn't print.
// Used to avoid leaking credentials when logging.
type SecureString string

// String format.
func (s SecureString) String() string {
	return "[REDACTED]"
}

// MarshalJSON format.
func (s SecureString) MarshalJSON() ([]byte, error) {
	return json.Marshal("[REDACTED]") //nolint:wrapcheck
}

var errConversionBody = errors.New("error converting request body")

// helper struct to have a json reader object.
type jsonBody struct {
	data       any
	marshaller func(any) ([]byte, error)
	reader     func([]byte, []byte) (int, error)
}

func newJSONBody(data any) *jsonBody {
	return &jsonBody{
		data:       data,
		marshaller: json.Marshal,
		reader: func(data []byte, buf []byte) (int, error) {
			return bytes.NewReader(data).Read(buf)
		},
	}
}

// Read as json.
func (b jsonBody) Read(buf []byte) (int, error) {
	jsonData, err := b.marshaller(b.data)
	if err != nil {
		return 0, errors.Join(errConversionBody, err)
	}

	read, err := b.reader(jsonData, buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return read, errors.Join(errConversionBody, err)
	}

	return read, err
}
