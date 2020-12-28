package jsonrpc

import (
	"context"
	"encoding/json"
)

// An endpoint that implements this can serve this transport
type JSONRPCServer interface {
	JSONRPCServerRequestDecoder
	JSONRPCServerResponseEncoder
}

//
type JSONRPCServerRequestDecoder interface {
	DecodeJSONRPCRequest(context.Context, json.RawMessage) (interface{}, error)
}

//
type JSONRPCServerResponseEncoder interface {
	EncodeJSONRPCResponse(context.Context, interface{}) (json.RawMessage, error)
}
