package http

import (
	"github.com/gomatic/go-kit-phases/pkg/api/instance"
)

func Server() instance.Server { return instance.New("Moody", instance.HTTPTransport, Listen) }
