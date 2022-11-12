package server

import (
	"github.com/gomatic/go-kit-phases/internal/api/moody/transport/jsonrpc"
)

func init() {
	// TODO make configurable
	servers["127.0.0.1:8095"] = jsonrpc.Server()
}
