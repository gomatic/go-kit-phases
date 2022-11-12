package version

import (
	"github.com/urfave/cli/v2"
)

//
func action(c *cli.Context) error {
	return nil
}

//
func Commands() cli.Commands {
	return cli.Commands{
		{
			Name:   "version",
			Usage:  "Shows versions",
			Action: action,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "output as JSON",
				},
			},
		},
	}
}
