package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gomatic/go-kit-phases/api/moody"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	"github.com/gomatic/go-kit-phases/internal/api/service"
	feeling "github.com/gomatic/go-kit-phases/internal/api/transform/feeling/server/grpc"
	feelings "github.com/gomatic/go-kit-phases/internal/api/transform/feelings/server/grpc"
	overall "github.com/gomatic/go-kit-phases/internal/api/transform/overall/server/grpc"
	query "github.com/gomatic/go-kit-phases/internal/api/transform/query/server/grpc"
)

//
func NewServer(endpoints apiendpoints.Set) moody.SelfServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.Create.ToEndpoint(),
			feeling.Request,
			overall.Response,
		),
		retrieve: grpctransport.NewServer(
			endpoints.Retrieve.ToEndpoint(),
			query.Request,
			feeling.Response,
		),
		update: grpctransport.NewServer(
			endpoints.Update.ToEndpoint(),
			feeling.Request,
			overall.Response,
		),
		delete: grpctransport.NewServer(
			endpoints.Delete.ToEndpoint(),
			feeling.Request,
			feeling.Response,
		),
		list: grpctransport.NewServer(
			endpoints.List.ToEndpoint(),
			query.Request,
			feelings.Response,
		),
	}
}

//
type grpcServer struct {
	moody.UnimplementedSelfServer
	create   grpctransport.Handler
	retrieve grpctransport.Handler
	update   grpctransport.Handler
	delete   grpctransport.Handler
	list     grpctransport.Handler
}

//
func (s *grpcServer) Create(ctx context.Context, feelings *moody.Feeling) (*moody.Overall, error) {
	return service.Self{}.Create(ctx, feelings)
}

//
func (s *grpcServer) Retrieve(ctx context.Context, request *moody.Query) (*moody.Feeling, error) {
	return service.Self{}.Retrieve(ctx, request)
}

//
func (s *grpcServer) Update(ctx context.Context, feelings *moody.Feeling) (*moody.Overall, error) {
	return service.Self{}.Update(ctx, feelings)
}

//
func (s *grpcServer) Delete(ctx context.Context, feelings *moody.Feeling) (*moody.Feeling, error) {
	return service.Self{}.Delete(ctx, feelings)
}

//
func (s *grpcServer) List(ctx context.Context, request *moody.Query) (*moody.Feelings, error) {
	return service.Self{}.List(ctx, request)
}
