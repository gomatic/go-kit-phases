package server

import (
	"github.com/gomatic/go-kit-phases/internal/api/moody/transport/http"
)

func init() {
	// TODO make configurable
	servers["127.0.0.1:8093"] = http.Server()
}
