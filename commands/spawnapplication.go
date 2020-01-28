package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/flags"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/prompt"
	"github.com/elko-dev/spawn/selections"
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

func Temp(action SpawnAction) cli.Command {
	return cli.Command{
		Name:    "temp",
		Aliases: []string{"temp"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			command := prompt.UserCommands{}
			selection := prompt.Selection{command}
			userSelections, _ := selection.Application()
			executeActionTemp(action, userSelections)
			return nil
		},
	}
}
func promptUserForInput() (platform.Application, error) {
	//TODO: Consider refactoring to builder
	application := platform.Application{}

	_, applicationType, err := selections.ApplicationType()
	application.ApplicationType = applicationType

	if err != nil {
		println("Error selecting application type")
		return platform.Application{}, err
	}

	projectName, err := prompt.ProjectName()
	if err != nil {
		println("Invalid Project Name")
		return platform.Application{}, err
	}
	application.ProjectName = projectName

	useCustomTemplate, err := prompt.UseCustomTemplate()

	if strings.ToLower(useCustomTemplate) == "y" {
		templateURL, err := prompt.TemplateURL()
		if err != nil {
			println("Template URL Failed")
			return platform.Application{}, err
		}
		application.TemplateURL = templateURL
	}

	platformToken, err := prompt.PlatformToken()
	if err != nil {
		println("Invalid DeployToken")
		return platform.Application{}, err
	}
	application.PlatformToken = platformToken

	gitToken, err := prompt.GitlabAccessToken()
	if err != nil {
		println("Invalid Git Token")
		return platform.Application{}, err
	}
	application.GitToken = gitToken

	environments := []string{"dev", "stage", "prod"}
	application.Environments = environments

	return application, nil
}

func executeAction(action SpawnAction, application platform.Application) error {
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

	return nil
}

func executeActionTemp(action SpawnAction, userCommands prompt.UserSelections) {
	//create client
	clientApplication := createClientApplication(userCommands)
	executeApp(action, clientApplication)
	//create server
	serverApplication := createServerApplication(userCommands)
	executeApp(action, serverApplication)
}

func executeApp(action SpawnAction, application platform.Application) {
	fmt.Printf("%+v\n", application)

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
	clientApplication.ProjectName = "HardCode"
	clientApplication.Environments = []string{"dev", "stage", "prod"}
	clientApplication.ApplicationType = userCommands.ClientLanguageType
	clientApplication.ProjectName = userCommands.ProjectName + "-client"
	//TODO: factor this out
	gitToken, _ := prompt.GitlabAccessToken()
	clientApplication.GitToken = gitToken

	platformToken, _ := prompt.PlatformToken()
	clientApplication.PlatformToken = platformToken
	return clientApplication
}

func createServerApplication(userCommands prompt.UserSelections) platform.Application {
	serverApplication := platform.Application{}
	serverApplication.ProjectName = "HardCode"
	serverApplication.Environments = []string{"dev", "stage", "prod"}
	serverApplication.ApplicationType = userCommands.ServerType
	serverApplication.ProjectName = userCommands.ProjectName + "-server"

	//TODO: factor this out
	gitToken, _ := prompt.GitlabAccessToken()
	serverApplication.GitToken = gitToken

	platformToken, _ := prompt.PlatformToken()
	serverApplication.PlatformToken = platformToken
	return serverApplication
}
