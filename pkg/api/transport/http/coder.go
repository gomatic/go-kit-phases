package http

import (
	"context"
	"net/http"
)

// Coder allows an endpoint implementation to transform messages.
type Coder interface {
	ServerRequestDecoder
	ServerResponseEncoder
}

// ServerRequestDecoder decodes an HTTP message.
type ServerRequestDecoder interface {
	DecodeHTTPRequest(context.Context, *http.Request) (interface{}, error)
}

// ServerResponseEncoder encodes an HTTP message.
type ServerResponseEncoder interface {
	EncodeHTTPResponse(context.Context, http.ResponseWriter, interface{}) error
}
