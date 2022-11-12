package instance

import (
	"fmt"
	stdlog "log"
	"sort"
	"strings"

	"github.com/go-kit/kit/log"
)

type Server interface {
	ServiceName() string
	TransportName() transportName
	Listener() ListenerFunc
	Interrupt(error)
	Executor(address Address, logger log.Logger) func() error
}

type Address string

type transportName string

const (
	JSONRPCTransport transportName = "JSONRPC"
	GRPCTransport    transportName = "gRPC"
	HTTPTransport    transportName = "HTTP"
)

type Servers map[Address]Server

func (s Servers) String() string {
	var ss sort.StringSlice
	for a, l := range s {
		ss = append(ss, fmt.Sprintf("%s %s/%s", a, l.ServiceName(), l.TransportName()))
	}
	ss.Sort()
	return strings.Join(ss, "\n")
}

var _ fmt.Stringer = Servers{}

func (s Servers) ByTransport(name transportName) (ss Servers) {
	ss = Servers{}
	for a, l := range s {
		if l.TransportName() == name {
			ss[a] = l
		}
	}
	return ss
}

type ListenerFunc func(string, log.Logger) error

type instance struct {
	serviceName   string
	transportName transportName
	listener      ListenerFunc
}

func (i instance) ServiceName() string          { return i.serviceName }
func (i instance) TransportName() transportName { return i.transportName }
func (i instance) Listener() ListenerFunc       { return i.listener }
func (i instance) Interrupt(err error)          {}

//
func (i instance) Executor(address Address, logger log.Logger) func() error {
	return func() error {
		if err := logger.Log("service", i.serviceName, "transport", i.transportName, "addr", address); err != nil {
			stdlog.Println(err)
		}
		if i.listener == nil {
			return fmt.Errorf("%w: %s/%s/%s", ErrorNoListener, i.serviceName, i.transportName, address)
		}
		return i.listener(string(address), logger)
	}
}

var _ Server = instance{}

func New(serviceName string, transportName transportName, listener ListenerFunc) Server {
	return instance{
		serviceName:   serviceName,
		transportName: transportName,
		listener:      listener,
	}
}
