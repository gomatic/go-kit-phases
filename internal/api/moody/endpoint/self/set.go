package self

import (
	"github.com/go-kit/kit/endpoint"
	api "github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/create"
	"github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/delete"
	"github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/list"
	"github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/retrieve"
	"github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/update"
)

// Set is the methods for the endpoint
type Set struct {
	Create   create.Endpoint
	Retrieve retrieve.Endpoint
	Update   update.Endpoint
	Delete   delete.Endpoint
	List     list.Endpoint
}

// New creates a Set for the service
func New(svc api.SelfServer) Set {
	return Set{
		Create:   create.New(svc),
		Retrieve: retrieve.New(svc),
		Update:   update.New(svc),
		Delete:   delete.New(svc),
		List:     list.New(svc),
	}
}

// EndpointsSlice creates a slice of service endpoints
func (s Set) EndpointsSlice() (es []endpoint.Endpoint) {
	es = append(es, s.Create.ToEndpoint())
	es = append(es, s.Retrieve.ToEndpoint())
	es = append(es, s.Update.ToEndpoint())
	es = append(es, s.Delete.ToEndpoint())
	es = append(es, s.List.ToEndpoint())
	return es
}

// EndpointsMap creates a map of service endpoints.
func (s Set) EndpointsMap() (es map[string]endpoint.Endpoint) {
	es = map[string]endpoint.Endpoint{
		"create":   s.Create.ToEndpoint(),
		"retrieve": s.Retrieve.ToEndpoint(),
		"update":   s.Update.ToEndpoint(),
		"delete":   s.Delete.ToEndpoint(),
		"list":     s.List.ToEndpoint(),
	}
	return es
}
