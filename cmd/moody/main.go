package main

import (
	"fmt"
	"os"
	"sort"

	cmd_client "github.com/gomatic/go-kit-phases/internal/commands/client"
	cmd_server "github.com/gomatic/go-kit-phases/internal/commands/server"
	cmd_version "github.com/gomatic/go-kit-phases/internal/commands/version"
	"github.com/gomatic/go-kit-phases/pkg/commander"
	"github.com/urfave/cli/v2"
)

// main is the entrypoint for the command.
func main() { _ = run(os.Args) }

// run is separated from main to make it more straightforward to write unit tests for the command.
// To unit test this function,
func run(args []string) error {
	app := &cli.App{
		Commands: commander.New(
			cmd_client.Commands,
			cmd_server.Commands,
			cmd_version.Commands,
		),
		Name:                   "moody",
		HelpName:               "",
		Usage:                  "moody",
		UsageText:              "moody [options] (client|server) [args]...",
		Description:            "moody",
		EnableBashCompletion:   true,
		HideHelp:               false,
		HideVersion:            false,
		BashComplete:           nil,
		CommandNotFound:        commandNotFound,
		OnUsageError:           onUsageError,
		Writer:                 os.Stderr,
		ErrWriter:              os.Stderr,
		ExitErrHandler:         exitErrHandler,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: true,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app.Run(args)
}

func exitErrHandler(c *cli.Context, err error) {
	if err != nil {
		code := 1
		defer func() { _, _ = fmt.Fprint(os.Stderr, err) }()
		switch coder := err.(type) {
		case cli.ExitCoder:
			code = coder.ExitCode()
		}
		if code == 0 {
			code = 1
		}
		_ = cli.ShowSubcommandHelp(c)
		cli.HandleExitCoder(cli.Exit(err, code))
	}
}

//
func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	if isSubcommand {
		_ = cli.ShowSubcommandHelp(c)
		return err
	}

	_, _ = fmt.Fprintln(c.App.Writer, err)
	return nil
}

//
func commandNotFound(ctx *cli.Context, command string) {
	_, _ = fmt.Fprintf(ctx.App.Writer, "Not a valid command: %s\n", command)
	_ = cli.ShowAppHelp(ctx)
}
