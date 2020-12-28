package jsonrpc

import (
	"net"
	"net/http"

	apiendpoints "github.com/gomatic/go-kit-phases/internal/api/endpoint"
)

//
func Listen(address string, endpoints apiendpoints.Set) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return http.Serve(listener, NewServer(endpoints))
}
