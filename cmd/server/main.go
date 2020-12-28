package main

import (
	stdlog "log"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/gomatic/go-kit-phases/internal/api/endpoint"
	"github.com/gomatic/go-kit-phases/internal/api/server"
	"github.com/gomatic/go-kit-phases/internal/api/service"
	"github.com/oklog/oklog/pkg/group"
)

//
func main() {
	run(os.Args)
}

//
var servers = map[string]server.Instance{}

//
func run(_ []string) {
	endpoints := endpoint.New(service.New())

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var g group.Group
	for address, svr := range servers {
		svr := svr
		g.Add(svr.Executor(address, endpoints, logger), svr.Interrupt)
	}

	if err := logger.Log("exit", g.Run()); err != nil {
		stdlog.Println(err)
	}
}
