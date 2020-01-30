package commands

import (
	"os"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/flags"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/prompt"
	"github.com/urfave/cli"
)

// SpawnAction describing the functionality to Create repositories
type SpawnAction interface {
	Application(app applications.App, application platform.Application) error
}

// Run is the method to run the CreateRepository command
func Run(action SpawnAction) cli.Command {
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

func executeAction(action SpawnAction, userCommands prompt.UserSelections) {
	//create client
	clientApplication := createClientApplication(userCommands)
	createApp(action, clientApplication)
	//create server
	serverApplication := createServerApplication(userCommands)
	createApp(action, serverApplication)
}

func createApp(action SpawnAction, application platform.Application) {
	app, err := applications.CreateApp(application)
	if err != nil {
		println("Error creating application.  Please verify your parameters are correct or submit an issue to Github")
		os.Exit(1)
	}
	err = action.Application(app, application)
	if err != nil {
		println("Some number of operations failed, exiting...")
		os.Exit(1)
	}
}

func createClientApplication(userCommands prompt.UserSelections) platform.Application {
	clientApplication := platform.Application{}
	clientApplication.Environments = []string{"dev", "stage", "prod"}
	clientApplication.ApplicationType = userCommands.ClientLanguageType
	clientApplication.ProjectName = userCommands.ProjectName + "-client"
	clientApplication.GitToken = userCommands.GitToken
	clientApplication.PlatformToken = userCommands.PlatformToken
	clientApplication.PlatformTeamName = userCommands.PlatformTeamName
	return clientApplication
}

func createServerApplication(userCommands prompt.UserSelections) platform.Application {
	serverApplication := platform.Application{}
	serverApplication.Environments = []string{"dev", "stage", "prod"}
	serverApplication.ApplicationType = userCommands.ServerType
	serverApplication.ProjectName = userCommands.ProjectName + "-server"
	serverApplication.GitToken = userCommands.GitToken
	serverApplication.PlatformToken = userCommands.PlatformToken
	serverApplication.PlatformTeamName = userCommands.PlatformTeamName
	return serverApplication
}
