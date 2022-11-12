package self

import (
	"github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	transport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/gomatic/go-kit-phases/api/moody"
	api "github.com/gomatic/go-kit-phases/api/moody"
	endpint "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	create "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/create/grpc"
	delete "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/delete/grpc"
	list "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/list/grpc"
	retrieve "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/retrieve/grpc"
	update "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/update/grpc"
)

//
func New(e endpint.Set, logger log.Logger) api.SelfServer {
	options := []transport.ServerOption{
		transport.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		transport.ServerBefore(jwt.GRPCToContext()),
	}

	return &serviceInstance{
		create: transport.NewServer(
			e.Create.ToEndpoint(),
			create.Request,
			create.Response,
			options...,
		),
		retrieve: transport.NewServer(
			e.Retrieve.ToEndpoint(),
			retrieve.Request,
			retrieve.Response,
			options...,
		),
		update: transport.NewServer(
			e.Update.ToEndpoint(),
			update.Request,
			update.Response,
			options...,
		),
		delete: transport.NewServer(
			e.Delete.ToEndpoint(),
			delete.Request,
			delete.Response,
			options...,
		),
		list: transport.NewServer(
			e.List.ToEndpoint(),
			list.Request,
			list.Response,
			options...,
		),
	}
}

//
type serviceInstance struct {
	moody.UnimplementedSelfServer
	create   transport.Handler
	retrieve transport.Handler
	update   transport.Handler
	delete   transport.Handler
	list     transport.Handler
}
