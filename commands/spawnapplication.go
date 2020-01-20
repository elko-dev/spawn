package commands

import (
	"os"
	"strings"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/flags"
	"github.com/elko-dev/spawn/prompt"
	"github.com/elko-dev/spawn/selections"
	"github.com/urfave/cli"
)

// SpawnAction describing the functionality to Create repositories
type SpawnAction interface {
	Application(app applications.App, environments []string) error
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

func promptUserForInput() (applications.Application, error) {
	//TODO: Consider refactoring to builder
	application := applications.Application{}

	_, applicationType, err := selections.ApplicationType()
	application.ApplicationType = applicationType

	if err != nil {
		println("Error selecting application type")
		return applications.Application{}, err
	}

	projectName, err := prompt.ProjectName()
	if err != nil {
		println("Invalid Project Name")
		return applications.Application{}, err
	}
	application.ProjectName = projectName

	useCustomTemplate, err := prompt.UseCustomTemplate()
	if err != nil {
		println("Use Custom Template Failed")
		return applications.Application{}, err
	}

	if strings.ToLower(useCustomTemplate) == "y" {
		templateURL, err := prompt.TemplateURL()
		if err != nil {
			println("Template URL Failed")
			return applications.Application{}, err
		}
		application.TemplateURL = templateURL
	}

	deployToken, err := prompt.DeployAccessToken()
	if err != nil {
		println("Invalid DeployToken")
		return applications.Application{}, err
	}
	application.DeployToken = deployToken

	accessToken, err := prompt.GitlabAccessToken()
	if err != nil {
		println("Invalid AccessToken")
		return applications.Application{}, err
	}
	application.AccessToken = accessToken

	environments := []string{"dev", "stage", "prod"}
	application.Environments = environments

	return application, nil
}

func executeAction(action SpawnAction, application applications.Application) error {
	app, err := applications.CreateApp(application)
	if err != nil {
		println("Error creating application.  Please verify your parameters are correct or submit an issue to Github")
		os.Exit(1)
	}
	err = action.Application(app, application.Environments)
	if err != nil {
		println("Some number of operations failed, exiting...")
		os.Exit(1)
	}

	return nil
}
