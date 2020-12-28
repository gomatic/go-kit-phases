package http

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	feeling "github.com/gomatic/go-kit-phases/internal/api/transform/feeling/server/http"
	feelings "github.com/gomatic/go-kit-phases/internal/api/transform/feelings/server/http"
	overall "github.com/gomatic/go-kit-phases/internal/api/transform/overall/server/http"
	query "github.com/gomatic/go-kit-phases/internal/api/transform/query/server/http"
	"github.com/gorilla/mux"
)

//
func NewServer(endpoints apiendpoints.Set) http.Handler {
	m := mux.NewRouter()
	r := m.PathPrefix("/moody").Subrouter()
	r.Methods("POST").Handler(httptransport.NewServer(
		endpoints.Create.ToEndpoint(),
		feeling.Request,
		overall.Response,
	))
	r.Methods("GET").Handler(httptransport.NewServer(
		endpoints.Retrieve.ToEndpoint(),
		query.Request,
		feeling.Response,
	))
	r.Methods("PUT").Handler(httptransport.NewServer(
		endpoints.Update.ToEndpoint(),
		feeling.Request,
		overall.Response,
	))
	r.Methods("DELETE").Handler(httptransport.NewServer(
		endpoints.Delete.ToEndpoint(),
		feeling.Request,
		feeling.Response,
	))
	r.Methods("GET").Path("/list").Handler(httptransport.NewServer(
		endpoints.List.ToEndpoint(),
		query.Request,
		feelings.Response,
	))
	return m
}
