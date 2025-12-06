// Package trading212 github.com/cyrbil/go-trading212
package trading212

import (
	"bytes"
	"encoding/json"
	"errors"
	"iter"
)

var (
	errRequestNil       = errors.New("request is nil")
	errRequestEmpty     = errors.New("request is empty")
	errEmptyIter        = errors.New("error, empty iterator")
	errDecodingResponse = errors.New("error reading response json")
)

// Response is a future abstraction that will normalize any API response
// (object, array or paginated array). Implement detection/parsing in a follow-up change.
type Response[T any] struct {
	err     error
	request IRequest
	raw     *json.RawMessage
}

// paginatedResponse is a generic wrapper for paginated API responses.
type paginatedResponse struct {
	Items        *json.RawMessage `json:"items"`
	NextPagePath *string          `json:"nextPagePath"`
}

func (r *Response[T]) validate() error {
	if r.err != nil {
		return r.err
	}

	if r.request == nil {
		return errRequestNil
	}

	if r.raw == nil {
		return errRequestEmpty
	}

	return nil
}

// Object get single response object
func (r *Response[T]) Object() (*T, error) {
	err := r.validate()
	if err != nil {
		return nil, err
	}

	iterator, err := r.Items()
	if err != nil {
		return nil, err
	}

	next, stop := iter.Pull(iterator)
	defer stop()

	value, ok := next()
	if !ok {
		return nil, errEmptyIter
	}

	return value, nil
}

// Items get iterator over response results
//
//nolint:cyclop,funlen  // TODO: too complex
func (r *Response[T]) Items() (iter.Seq[*T], error) {
	err := r.validate()
	if err != nil {
		return nil, err
	}

	// detect paginated results
	var paginatedResponse paginatedResponse
	decoder := json.NewDecoder(bytes.NewBuffer(*r.raw)) //nolint:wsl,wsl_v5
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&paginatedResponse)
	if err != nil {
		// assume data is array, but use like paginated
		paginatedResponse.Items = r.raw
		paginatedResponse.NextPagePath = nil
	}

	// decode current array
	var data []*T
	decoder = json.NewDecoder(bytes.NewBuffer(*paginatedResponse.Items)) //nolint:wsl,wsl_v5
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&data)
	if err != nil {
		if len(data) > 0 {
			return nil, errors.Join(errDecodingResponse, err)
		}
		var value *T
		decoder = json.NewDecoder(bytes.NewBuffer(*paginatedResponse.Items)) //nolint:wsl,wsl_v5
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&value)
		if err != nil {
			return nil, errors.Join(errDecodingResponse, err)
		}

		data = []*T{value}
	}

	iterator := func(yield func(*T) bool) {
		for _, value := range data {
			if !yield(value) {
				return
			}
		}

		if paginatedResponse.NextPagePath == nil || *paginatedResponse.NextPagePath == "" {
			return
		}

		httpRequest := r.request.http()
		query := httpRequest.URL.Query()
		query.Set("cursor", *paginatedResponse.NextPagePath)
		httpRequest.URL.RawQuery = query.Encode()

		data, err := r.request.Do()
		if err != nil {
			return
		}

		response := &Response[T]{request: r.request, raw: data, err: nil}

		nextIterator, err := response.Items()
		if err != nil {
			return
		}

		next, stop := iter.Pull(nextIterator)
		defer stop()

		for {
			value, ok := next()
			if !ok {
				return
			}

			if !yield(value) {
				return
			}
		}
	}

	return iterator, nil
}
