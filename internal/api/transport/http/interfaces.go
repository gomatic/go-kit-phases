package http

import (
	"context"
	"net/http"
)

// An endpoint that implements this can serve this transport
type HTTPServer interface {
	HTTPServerRequestDecoder
	HTTPServerResponseEncoder
}

//
type HTTPServerRequestDecoder interface {
	DecodeHTTPRequest(context.Context, *http.Request) (interface{}, error)
}

//
type HTTPServerResponseEncoder interface {
	EncodeHTTPResponse(context.Context, http.ResponseWriter, interface{}) error
}
