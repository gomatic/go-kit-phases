package jsonrpc

import (
	"github.com/go-kit/kit/transport/http/jsonrpc"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	feeling "github.com/gomatic/go-kit-phases/internal/api/transform/feeling/server/jsonrpc"
	feelings "github.com/gomatic/go-kit-phases/internal/api/transform/feelings/server/jsonrpc"
	overall "github.com/gomatic/go-kit-phases/internal/api/transform/overall/server/jsonrpc"
	query "github.com/gomatic/go-kit-phases/internal/api/transform/query/server/jsonrpc"
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
			Endpoint: endpoints.Create.ToEndpoint(),
			Decode:   feeling.Request,
			Encode:   overall.Response,
		},
		"retrieve": jsonrpc.EndpointCodec{
			Endpoint: endpoints.Retrieve.ToEndpoint(),
			Decode:   query.Request,
			Encode:   feeling.Response,
		},
		"update": jsonrpc.EndpointCodec{
			Endpoint: endpoints.Update.ToEndpoint(),
			Decode:   feeling.Request,
			Encode:   overall.Response,
		},
		"delete": jsonrpc.EndpointCodec{
			Endpoint: endpoints.Delete.ToEndpoint(),
			Decode:   feeling.Request,
			Encode:   feeling.Response,
		},
		"list": jsonrpc.EndpointCodec{
			Endpoint: endpoints.List.ToEndpoint(),
			Decode:   feeling.Request,
			Encode:   feelings.Response,
		},
	}
}
