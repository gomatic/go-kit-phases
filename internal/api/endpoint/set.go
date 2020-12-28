package endpoint

import (
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
func (s Set) EndpointsSlice() (es []endpoint.Endpoint) {
	es = append(es, s.CreateEndpoint)
	es = append(es, s.RetrieveEndpoint)
	es = append(es, s.UpdateEndpoint)
	es = append(es, s.DeleteEndpoint)
	es = append(es, s.ListEndpoint)
	return es
}

//
func (s Set) EndpointsMap() (es map[string]endpoint.Endpoint) {
	es = map[string]endpoint.Endpoint{
		"create":   s.CreateEndpoint,
		"retrieve": s.RetrieveEndpoint,
		"update":   s.UpdateEndpoint,
		"delete":   s.DeleteEndpoint,
		"list":     s.ListEndpoint,
	}
	return es
}

//
func New(svc moody.SelfServer) Set {
	return Set{
		CreateEndpoint:   NewCreateEndpoint(svc),
		RetrieveEndpoint: NewRetrieveEndpoint(svc),
		UpdateEndpoint:   NewUpdateEndpoint(svc),
		DeleteEndpoint:   NewDeleteEndpoint(svc),
		ListEndpoint:     NewListEndpoint(svc),
	}
}
