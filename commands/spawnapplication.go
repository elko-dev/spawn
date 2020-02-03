package commands

import (
	"github.com/elko-dev/spawn/flags"
	"github.com/elko-dev/spawn/prompt"
	"github.com/urfave/cli"
)

type ApplicationType interface {
	Create() error
}

// Run is the method to run the CreateRepository command
func Run() cli.Command {
	return cli.Command{
		Name:    "application",
		Aliases: []string{"application"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			command := prompt.UserCommands{}
			platform := prompt.HerokuCommand{}
			git := prompt.GitPrompts{}
			selection := prompt.Selection{command, platform, git}
			selection.Application()
			return nil
		},
	}
}
