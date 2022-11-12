package jsonrpc

import (
	"net"
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	transport "github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/go-kit/log"
	selfendpoint "github.com/gomatic/go-kit-phases/internal/api/moody/endpoint/self"
	selfservice "github.com/gomatic/go-kit-phases/internal/api/moody/service/self"
	selfserver "github.com/gomatic/go-kit-phases/internal/api/moody/transport/jsonrpc/self"
)

//
func Listen(address string, logger log.Logger) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	options := []transport.ServerOption{
		transport.ServerErrorLogger(logger),
		transport.ServerBefore(jwt.HTTPToContext()),
	}

	endpoints := map[string]transport.EndpointCodecMap{
		"self": selfserver.New(selfendpoint.New(selfservice.New())),
		// Template: "{{.Service}}":    {{.Service}}server.New({{.Service}}endpoint.New({{.Service}}service.New())),
	}

	endpointMap := transport.EndpointCodecMap{}
	for name, em := range endpoints {
		for k, v := range em {
			endpointMap[name+"/"+k] = v
		}
	}

	return http.Serve(listener, transport.NewServer(endpointMap, options...))
}
