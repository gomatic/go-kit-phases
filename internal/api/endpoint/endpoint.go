package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gomatic/go-kit-phases/api/moody"
)

//
type Set struct {
	CreateEndpoint   endpoint.Endpoint
	RetrieveEndpoint endpoint.Endpoint
	UpdateEndpoint   endpoint.Endpoint
	DeleteEndpoint   endpoint.Endpoint
	ListEndpoint     endpoint.Endpoint
}

//
func New(svc moody.SelfServer) Set {
	return Set{
		CreateEndpoint:     NewCreateEndpoint(svc),
		RetrieveEndpoint:   NewRetrieveEndpoint(svc),
		UpdateEndpoint:     NewUpdateEndpoint(svc),
		DeleteEndpoint:     NewDeleteEndpoint(svc),
		ListEndpoint:       NewListEndpoint(svc),
	}
}

//
func NewCreateEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("implement me")
	}
}

//
func NewRetrieveEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("implement me")
	}
}

//
func NewUpdateEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("implement me")
	}
}

//
func NewDeleteEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("implement me")
	}
}

//
func NewListEndpoint(s moody.SelfServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		panic("implement me")
	}
}
