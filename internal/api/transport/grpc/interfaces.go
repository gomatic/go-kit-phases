package grpc

import (
	"context"
)

// An endpoint that implements this can serve this transport
type GRPCServer interface {
	GRPCServerRequestDecoder
	GRPCServerResponseEncoder
}

//
type GRPCServerRequestDecoder interface {
	DecodeGRPCRequest(context.Context, interface{}) (interface{}, error)
}

//
type GRPCServerResponseEncoder interface {
	EncodeGRPCResponse(context.Context, interface{}) (interface{}, error)
}
