package http

import (
	"github.com/gomatic/go-kit-phases/internal/api/server"
)

//
func Server() server.Instance {
	return server.Instance{
		Transport: "HTTP",
		Listener:  Listen,
	}
}
