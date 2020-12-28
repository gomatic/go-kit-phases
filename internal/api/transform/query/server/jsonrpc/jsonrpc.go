package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/transform"
	"github.com/gomatic/go-kit-phases/internal/api/transform/feeling/server"
)

//
func Request(_ context.Context, request json.RawMessage) (interface{}, error) {
	return FromJSON(request)
}

//
func FromJSON(request json.RawMessage) (interface{}, error) {
	var value *moody.Query
	return value, json.Unmarshal(request, &value)
}

//
func Response(_ context.Context, response interface{}) (json.RawMessage, error) {
	return ToJSON(response)
}

//
func ToJSON(response interface{}) (json.RawMessage, error) {
	return transform.Marshaller(server.FromInterface(response))
}
