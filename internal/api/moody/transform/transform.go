package transform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gomatic/go-kit-phases/pkg/errors"
)

//
func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errors.ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

//
func Marshaller(i interface{}, err error) (json.RawMessage, error) {
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal %T: %w", i, err)
	}
	return b, nil
}
