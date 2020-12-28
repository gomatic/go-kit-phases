package main

import (
	"github.com/gomatic/go-kit-phases/internal/api/transport/http"
)

func init() {
	servers["127.0.0.1:8081"] = http.Server()
}
