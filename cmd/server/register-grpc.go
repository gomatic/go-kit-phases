package main

import (
	"github.com/gomatic/go-kit-phases/internal/api/transport/grpc"
)

func init() {
	servers["127.0.0.1:8082"] = grpc.Server()
}
