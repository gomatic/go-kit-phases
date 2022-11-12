package self

import (
	transport "github.com/go-kit/kit/transport/http/jsonrpc"
	endpoints "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	create "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/create/jsonrpc"
	delete "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/delete/jsonrpc"
	list "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/list/jsonrpc"
	retrieve "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/retrieve/jsonrpc"
	update "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/update/jsonrpc"
)

//
func New(e endpoints.Set) transport.EndpointCodecMap {
	return transport.EndpointCodecMap{
		"create": transport.EndpointCodec{
			Endpoint: e.Create.ToEndpoint(),
			Decode:   create.Request,
			Encode:   create.Response,
		},
		"retrieve": transport.EndpointCodec{
			Endpoint: e.Retrieve.ToEndpoint(),
			Decode:   retrieve.Request,
			Encode:   retrieve.Response,
		},
		"update": transport.EndpointCodec{
			Endpoint: e.Update.ToEndpoint(),
			Decode:   update.Request,
			Encode:   update.Response,
		},
		"delete": transport.EndpointCodec{
			Endpoint: e.Delete.ToEndpoint(),
			Decode:   delete.Request,
			Encode:   delete.Response,
		},
		"list": transport.EndpointCodec{
			Endpoint: e.List.ToEndpoint(),
			Decode:   list.Request,
			Encode:   list.Response,
		},
	}
}
