package grpc

import (
	"context"
)

// Coder allows an endpoint implementation to transform messages.
type Coder interface {
	ServerRequestDecoder
	ServerResponseEncoder
}

// ServerRequestDecoder decodes a gRPC message.
type ServerRequestDecoder interface {
	DecodeGRPCRequest(context.Context, interface{}) (interface{}, error)
}

// ServerResponseEncoder encodes a gRPC message.
type ServerResponseEncoder interface {
	EncodeGRPCResponse(context.Context, interface{}) (interface{}, error)
}
