package grpc

import (
	"net"

	transport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	api "github.com/gomatic/go-kit-phases/api/moody"
	selfendpoint "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	selfservice "github.com/gomatic/go-kit-phases/internal/api/moody/service/self"
	selftransport "github.com/gomatic/go-kit-phases/internal/api/moody/transport/grpc/self"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//
func Listen(address string, logger log.Logger) error {
	grpcListener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(transport.Interceptor))
	reflection.Register(baseServer)

	api.RegisterSelfServer(baseServer, selftransport.New(selfendpoint.New(selfservice.New()), logger))
	// Template: api.Register{{.Service}}Server(baseServer, {{.Service}}transport.New({{.Service}}endpoint.New({{.Service}}service.New()), logger))

	return baseServer.Serve(grpcListener)
}
