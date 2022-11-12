package http

import (
	"net"
	"net/http"

	"github.com/go-kit/log"
	selfendpoint "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	selfservice "github.com/gomatic/go-kit-phases/internal/api/moody/service/self"
	selftransport "github.com/gomatic/go-kit-phases/internal/api/moody/transport/http/self"
	"github.com/gorilla/mux"
)

//
func Listen(address string, logger log.Logger) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	m := mux.NewRouter()
	selftransport.New(m, selfendpoint.New(selfservice.New()), logger)
	// Template: {{.Service}}transport.New(m, {{.Service}}endpoint.New({{.Service}}service.New()), logger)

	return http.Serve(listener, m)
}
