package self

import (
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	kittransport "github.com/go-kit/kit/transport"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	endpoints "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	create "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/create/http"
	delete "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/delete/http"
	list "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/list/http"
	retrieve "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/retrieve/http"
	update "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self/update/http"
	"github.com/gorilla/mux"
)

//
func New(m *mux.Router, e endpoints.Set, logger log.Logger) http.Handler {
	options := []transport.ServerOption{
		transport.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		transport.ServerBefore(jwt.HTTPToContext()),
	}

	r := m.PathPrefix("/moody").Subrouter()
	r.Methods("POST").Handler(transport.NewServer(
		e.Create.ToEndpoint(),
		create.Request,
		create.Response,
		options...,
	))
	r.Methods("GET").Handler(transport.NewServer(
		e.Retrieve.ToEndpoint(),
		retrieve.Request,
		retrieve.Response,
		options...,
	))
	r.Methods("PUT").Handler(transport.NewServer(
		e.Update.ToEndpoint(),
		update.Request,
		update.Response,
		options...,
	))
	r.Methods("DELETE").Handler(transport.NewServer(
		e.Delete.ToEndpoint(),
		delete.Request,
		delete.Response,
		options...,
	))
	r.Methods("GET").Path("/list").Handler(transport.NewServer(
		e.List.ToEndpoint(),
		list.Request,
		list.Response,
		options...,
	))
	return m
}
