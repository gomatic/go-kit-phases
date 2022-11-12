package http

import (
	"context"
	"net/http"

	in "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/moody/transform"
)

func Request(ctx context.Context, q *http.Request) (interface{}, error) {
	return &in.Feeling{}, nil
}
func Response(ctx context.Context, w http.ResponseWriter, p interface{}) error {
	return transform.EncodeHTTPGenericResponse(ctx, w, p)
}
