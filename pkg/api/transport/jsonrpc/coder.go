package jsonrpc

import (
	"context"
	"encoding/json"
)

// Coder allows an endpoint implementation to transform messages.
type Coder interface {
	ServerRequestDecoder
	ServerResponseEncoder
}

// ServerRequestDecoder decodes JSONRPC message.
type ServerRequestDecoder interface {
	DecodeJSONRPCRequest(context.Context, json.RawMessage) (interface{}, error)
}

// ServerResponseEncoder encodes JSONRPC message.
type ServerResponseEncoder interface {
	EncodeJSONRPCResponse(context.Context, interface{}) (json.RawMessage, error)
}
