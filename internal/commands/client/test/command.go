package tester

import (
	"github.com/gomatic/go-kit-phases/internal/commands/client/test/token"
	"github.com/gomatic/go-kit-phases/pkg/commander"
	"github.com/urfave/cli/v2"
)

//
func Commands() cli.Commands {
	return cli.Commands{
		{
			Name:        "test",
			Usage:       "client test",
			Description: "client test",
			Subcommands: commander.New(
				token.Commands,
			),
		},
	}
}
