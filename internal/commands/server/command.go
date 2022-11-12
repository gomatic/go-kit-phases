package server

import (
	stdlog "log"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/gomatic/go-kit-phases/pkg/api/instance"
	"github.com/oklog/oklog/pkg/group"
	"github.com/urfave/cli/v2"
)

//
var servers = instance.Servers{}

func Servers() instance.Servers {
	return servers
}

func Before(c *cli.Context) error {
	return nil
}

func Action(c *cli.Context) error {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var g group.Group
	for address, svr := range servers {
		svr := svr
		g.Add(svr.Executor(address, logger), svr.Interrupt)
	}

	if err := logger.Log("exit", g.Run()); err != nil {
		stdlog.Println(err)
	}

	return nil
}

func After(c *cli.Context) error {
	return nil
}

//
func Commands() cli.Commands {
	return cli.Commands{
		{
			Name:        "server",
			Aliases:     []string{"serve"},
			Usage:       "server",
			Description: "server",
			Action:      Action,
			Flags:       Flags(),
		},
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{}
}
