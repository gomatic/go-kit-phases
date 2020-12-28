package grpc

import (
	"context"

	"github.com/gomatic/go-kit-phases/internal/api/transform/feeling/server"
)

//
func Request(_ context.Context, request interface{}) (interface{}, error) {
	return server.FromInterface(request)
}

//
func Response(_ context.Context, response interface{}) (interface{}, error) {
	return server.FromInterface(response)
}
