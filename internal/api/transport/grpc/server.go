package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gomatic/go-kit-phases/api/moody"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	"github.com/gomatic/go-kit-phases/internal/api/service"
)

//
func NewServer(endpoints apiendpoints.Set) moody.SelfServer {
	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
		),
		retrieve: grpctransport.NewServer(
			endpoints.RetrieveEndpoint,
			decodeGRPCRetrieveRequest,
			encodeGRPCRetrieveResponse,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
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

//
func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return &moody.Feeling{}, nil
}

//
func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &moody.Overall{}, nil
}

//
func decodeGRPCRetrieveRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return &moody.Query{}, nil
}

//
func encodeGRPCRetrieveResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &moody.Feeling{}, nil
}

//
func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return &moody.Feeling{}, nil
}

//
func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &moody.Overall{}, nil
}

//
func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return &moody.Feeling{}, nil
}

//
func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &moody.Feeling{}, nil
}

//
func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return &moody.Query{}, nil
}

//
func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &moody.Feelings{}, nil
}
