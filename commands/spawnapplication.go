package commands

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/flags"
	"github.com/urfave/cli"
)

type ApplicationType interface {
	Create() error
}

// Run is the method to run the CreateRepository command
func Run(factory applicationtype.Factory) cli.Command {
	return cli.Command{
		Name:    "application",
		Aliases: []string{"application"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			return factory.CreateApplicationType().Create()
		},
	}
}
