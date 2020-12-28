package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
	"github.com/gorilla/mux"
)

//
func NewServer(endpoints apiendpoints.Set) http.Handler {
	m := mux.NewRouter()
	r := m.PathPrefix("/moody").Subrouter()
	r.Methods("POST").Handler(httptransport.NewServer(
		endpoints.CreateEndpoint,
		decodeHTTPCreateRequest,
		encodeHTTPGenericResponse,
	))
	r.Methods("GET").Handler(httptransport.NewServer(
		endpoints.RetrieveEndpoint,
		decodeHTTPRetrieveRequest,
		encodeHTTPGenericResponse,
	))
	r.Methods("PUT").Handler(httptransport.NewServer(
		endpoints.UpdateEndpoint,
		decodeHTTPUpdateRequest,
		encodeHTTPGenericResponse,
	))
	r.Methods("DELETE").Handler(httptransport.NewServer(
		endpoints.DeleteEndpoint,
		decodeHTTPDeleteRequest,
		encodeHTTPGenericResponse,
	))
	r.Methods("GET").Path("/list").Handler(httptransport.NewServer(
		endpoints.ListEndpoint,
		decodeHTTPListRequest,
		encodeHTTPGenericResponse,
	))
	return m
}

//
func decodeHTTPCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("implement me")
}

//
func decodeHTTPRetrieveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("implement me")
}

//
func decodeHTTPUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("implement me")
}

//
func decodeHTTPDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("implement me")
}

//
func decodeHTTPListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	panic("implement me")
}

//
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

//
type errorWrapper struct {
	Error string `json:"error"`
}

//
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
