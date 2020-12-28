package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/gomatic/go-kit-phases/api/moody"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint/create"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint/delete"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint/list"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint/retrieve"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint/update"
)

//
type Set struct {
	Create   create.Endpoint
	Retrieve retrieve.Endpoint
	Update   update.Endpoint
	Delete   delete.Endpoint
	List     list.Endpoint
}

//
func (s Set) EndpointsSlice() (es []endpoint.Endpoint) {
	es = append(es, endpoint.Endpoint(s.Create))
	es = append(es, endpoint.Endpoint(s.Retrieve))
	es = append(es, endpoint.Endpoint(s.Update))
	es = append(es, endpoint.Endpoint(s.Delete))
	es = append(es, endpoint.Endpoint(s.List))
	return es
}

//
func (s Set) EndpointsMap() (es map[string]endpoint.Endpoint) {
	es = map[string]endpoint.Endpoint{
		"create":   endpoint.Endpoint(s.Create),
		"retrieve": endpoint.Endpoint(s.Retrieve),
		"update":   endpoint.Endpoint(s.Update),
		"delete":   endpoint.Endpoint(s.Delete),
		"list":     endpoint.Endpoint(s.List),
	}
	return es
}

//
func New(svc moody.SelfServer) Set {
	return Set{
		Create:   create.New(svc),
		Retrieve: retrieve.New(svc),
		Update:   update.New(svc),
		Delete:   delete.New(svc),
		List:     list.New(svc),
	}
}
