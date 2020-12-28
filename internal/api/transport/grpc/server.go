package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/gomatic/go-kit-phases/api/moody"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
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
func (s *grpcServer) Create(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	panic("implement me")
}

//
func (s *grpcServer) Retrieve(_ context.Context, _ *moody.Query) (*moody.Feeling, error) {
	panic("implement me")
}

//
func (s *grpcServer) Update(_ context.Context, _ *moody.Feeling) (*moody.Overall, error) {
	panic("implement me")
}

//
func (s *grpcServer) Delete(_ context.Context, _ *moody.Feeling) (*moody.Feeling, error) {
	panic("implement me")
}

//
func (s *grpcServer) List(_ context.Context, _ *moody.Query) (*moody.Feelings, error) {
	panic("implement me")
}

//
func decodeGRPCCreateRequest(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func encodeGRPCCreateResponse(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func decodeGRPCRetrieveRequest(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func encodeGRPCRetrieveResponse(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func decodeGRPCUpdateRequest(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func encodeGRPCUpdateResponse(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func decodeGRPCDeleteRequest(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func encodeGRPCDeleteResponse(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func decodeGRPCListRequest(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

//
func encodeGRPCListResponse(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}
