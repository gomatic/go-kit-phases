package server

import (
	"github.com/gomatic/go-kit-phases/internal/api/moody/transport/grpc"
)

func init() {
	// TODO make configurable
	servers["127.0.0.1:8091"] = grpc.Server()
}
