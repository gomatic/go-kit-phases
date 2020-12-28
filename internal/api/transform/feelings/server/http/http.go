package http

import (
	"context"
	"net/http"

	"github.com/gomatic/go-kit-phases/internal/api/transform"
	"github.com/gomatic/go-kit-phases/api/moody"
)

//
func Request(ctx context.Context, request *http.Request) (interface{}, error) {
	return &moody.Feelings{}, nil
}

//
func Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return transform.EncodeHTTPGenericResponse(ctx, w, response)
}
