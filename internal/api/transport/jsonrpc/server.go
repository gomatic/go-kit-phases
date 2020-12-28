package jsonrpc

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/transport/http/jsonrpc"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
)

//
func NewServer(endpoints apiendpoints.Set) *jsonrpc.Server {
	handler := jsonrpc.NewServer(
		makeEndpointCodecMap(endpoints),
	)
	return handler
}

//
func makeEndpointCodecMap(endpoints apiendpoints.Set) jsonrpc.EndpointCodecMap {
	return jsonrpc.EndpointCodecMap{
		"create": jsonrpc.EndpointCodec{
			Endpoint: endpoints.CreateEndpoint,
			Decode:   decodeCreateRequest,
			Encode:   encodeCreateResponse,
		},
		"retrieve": jsonrpc.EndpointCodec{
			Endpoint: endpoints.RetrieveEndpoint,
			Decode:   decodeRetrieveRequest,
			Encode:   encodeRetrieveResponse,
		},
		"update": jsonrpc.EndpointCodec{
			Endpoint: endpoints.UpdateEndpoint,
			Decode:   decodeUpdateRequest,
			Encode:   encodeUpdateResponse,
		},
		"delete": jsonrpc.EndpointCodec{
			Endpoint: endpoints.DeleteEndpoint,
			Decode:   decodeDeleteRequest,
			Encode:   encodeDeleteResponse,
		},
		"list": jsonrpc.EndpointCodec{
			Endpoint: endpoints.ListEndpoint,
			Decode:   decodeListRequest,
			Encode:   encodeListResponse,
		},
	}
}

//
func decodeCreateRequest(_ context.Context, _ json.RawMessage) (interface{}, error) {
	panic("implement me")
}

//
func encodeCreateResponse(_ context.Context, _ interface{}) (json.RawMessage, error) {
	panic("implement me")
}

//
func decodeRetrieveRequest(_ context.Context, _ json.RawMessage) (interface{}, error) {
	panic("implement me")
}

//
func encodeRetrieveResponse(_ context.Context, _ interface{}) (json.RawMessage, error) {
	panic("implement me")
}

//
func decodeUpdateRequest(_ context.Context, _ json.RawMessage) (interface{}, error) {
	panic("implement me")
}

//
func encodeUpdateResponse(_ context.Context, _ interface{}) (json.RawMessage, error) {
	panic("implement me")
}

//
func decodeDeleteRequest(_ context.Context, _ json.RawMessage) (interface{}, error) {
	panic("implement me")
}

//
func encodeDeleteResponse(_ context.Context, _ interface{}) (json.RawMessage, error) {
	panic("implement me")
}

//
func decodeListRequest(_ context.Context, _ json.RawMessage) (interface{}, error) {
	panic("implement me")
}

//
func encodeListResponse(_ context.Context, _ interface{}) (json.RawMessage, error) {
	panic("implement me")
}
