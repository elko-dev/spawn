package commands

import (
	"os"

	"github.com/urfave/cli"
	"gitlab.com/shared-tool-chain/spawn/flags"
	"gitlab.com/shared-tool-chain/spawn/prompt"
)

// SpawnAction describing the functionality to Create repositories
type SpawnAction interface {
	Application(application Application) error
}

// Application is a struct representing a full application
type Application struct {
	ProjectName  string
	DeployToken  string
	AccessToken  string
	PlatformName string
	Environments []string
}

// Run is the method to run the CreateRepository command
func Run(action SpawnAction) cli.Command {
	return cli.Command{
		Name:    "application",
		Aliases: []string{"app"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			application, err := promptUserForInput()
			if err != nil {
				os.Exit(1)
			}
			return executeAction(action, application)
		},
	}
}

func promptUserForInput() (Application, error) {
	projectName, err := prompt.ProjectName()
	if err != nil {
		println("Invalid Project Name")
		return Application{}, err
	}

	platformTeamName, err := prompt.HerokuTeamName()
	if err != nil {
		println("Invalid Heroku Team Name")
		return Application{}, err
	}

	deployToken, err := prompt.DeployAccessToken()
	if err != nil {
		println("Invalid DeployToken")
		return Application{}, err
	}

	accessToken, err := prompt.GitlabAccessToken()
	if err != nil {
		println("Invalid AccessToken")
		return Application{}, err
	}
	environments := []string{"dev", "stage", "prod"}

	application := Application{
		ProjectName:  projectName,
		AccessToken:  accessToken,
		DeployToken:  deployToken,
		PlatformName: platformTeamName,
		Environments: environments,
	}
	return application, nil
}

func executeAction(action SpawnAction, application Application) error {
	err := action.Application(application)
	if err != nil {
		println("Some number of operations failed, exiting...")
		os.Exit(1)
	}

	return nil
}
