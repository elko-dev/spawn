package commands

import (
	"os"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/flags"
	"github.com/elko-dev/spawn/prompt"
	"github.com/elko-dev/spawn/web"
	"github.com/urfave/cli"
)

type ApplicationType interface {
	Create(action web.SpawnAction, userCommands prompt.UserSelections) error
}

// Run is the method to run the CreateRepository command
func Run(action web.SpawnAction) cli.Command {
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
			userSelections, _ := selection.Application()
			executeAction(action, userSelections)
			return nil
		},
	}
}

func executeAction(action web.SpawnAction, userCommands prompt.UserSelections) {
	var applicationType ApplicationType

	if userCommands.ApplicationType == constants.WebApplicationType {
		applicationType = web.WebType{}
	}

	if applicationType == nil {
		println("Unsupported Application Type")
		os.Exit(0)
	}
	err := applicationType.Create(action, userCommands)

	if err != nil {
		os.Exit(0)
	}
}
