package commander

import "github.com/urfave/cli/v2"

//
type CommandsFactory func() cli.Commands

//
type Commander interface {
	Commands() []cli.Command
}

//
func New(commands ...CommandsFactory) cli.Commands {
	cliCommands := make(cli.Commands, 0)
	for _, commandFactory := range commands {
		cliCommands = append(cliCommands, commandFactory()...)
	}
	return cliCommands
}
