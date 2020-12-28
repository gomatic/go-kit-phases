package grpc

import (
	"net"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/gomatic/go-kit-phases/api/moody"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	"google.golang.org/grpc"
)

//
func Listen(address string, endpoints apiendpoints.Set) error {
	grpcListener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := NewServer(endpoints)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	moody.RegisterSelfServer(baseServer, grpcServer)
	return baseServer.Serve(grpcListener)
}
