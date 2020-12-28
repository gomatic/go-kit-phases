package server

import (
	stdlog "log"

	"github.com/go-kit/kit/log"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint"
)

//
type Instance struct {
	Transport string
	Listener  func(string, endpoint.Set) error
}

//
func (s Instance) Executor(address string, endpoints endpoint.Set, logger log.Logger) func() error {
	return func() error {
		if err := logger.Log("transport", s.Transport, "addr", address); err != nil {
			stdlog.Println(err)
		}
		return s.Listener(address, endpoints)
	}
}

//
func (s Instance) Interrupt(err error) {}

