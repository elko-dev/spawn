package commands

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.com/shared-tool-chain/spawn/flags"
)

const (
	projectname = "projectname"
	deploytoken = "deploytoken"
	accesstoken = "accesstoken"
)

// SpawnAction describing the functionality to Create repositories
type SpawnAction interface {
	Application(application Application) error
}

// Application is a struct representing a full application
type Application struct {
	ProjectName string
	DeployToken string
	AccessToken string
}

// Run is the method to run the CreateRepository command
func Run(action SpawnAction) cli.Command {
	return cli.Command{
		Name:    "application",
		Aliases: []string{"app"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			application := Application{
				ProjectName: c.String(projectname),
				AccessToken: c.String(accesstoken),
				DeployToken: c.String(deploytoken),
			}
			return executeAction(action, application)
		},
	}
}

func executeAction(action SpawnAction, application Application) error {
	err := action.Application(application)
	if err != nil {
		println("Some number of operations failed, exiting...")
		os.Exit(1)
	}

	return nil
}
