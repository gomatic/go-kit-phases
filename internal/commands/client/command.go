package client

import (
	"fmt"

	tester "github.com/gomatic/go-kit-phases/internal/commands/client/test"
	"github.com/gomatic/go-kit-phases/internal/commands/server"
	"github.com/gomatic/go-kit-phases/pkg/commander"
	"github.com/urfave/cli/v2"
)

//
func action(c *cli.Context) error {
	fmt.Println(server.Servers())
	return nil
}

//
func Commands() cli.Commands {
	return cli.Commands{
		{
			Name:        "client",
			Aliases:     []string{"cli"},
			Usage:       "client",
			Description: "client",
			Action:      action,
			Flags:       []cli.Flag{},
			Subcommands: commander.New(
				tester.Commands,
			),
		},
	}
}
