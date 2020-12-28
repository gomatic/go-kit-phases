package main

import (
	"github.com/gomatic/go-kit-phases/internal/api/transport/jsonrpc"
)

func init() {
	servers["127.0.0.1:8084"] = jsonrpc.Server()
}
