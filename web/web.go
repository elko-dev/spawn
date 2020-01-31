package web

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/prompt"
)

// SpawnAction describing the functionality to Create repositories
type SpawnAction interface {
	Application(app applications.App, application platform.Application) error
}

// WebType struct to create an application type
type WebType struct {
	Client applications.App
	Server applications.App
}

// Create sets up a new application
func (webType WebType) Create(action SpawnAction, userCommands prompt.UserSelections) error {
	//create client
	clientApplication := createClientApplication(userCommands)
	err := createApp(action, clientApplication)
	if err != nil {
		return err
	}
	//create server
	serverApplication := createServerApplication(userCommands)
	return createApp(action, serverApplication)
}

func createApp(action SpawnAction, application platform.Application) error {
	app, err := applications.CreateApp(application)
	if err != nil {
		println("Error creating application.  Please verify your parameters are correct or submit an issue to Github")
		return err
	}
	err = action.Application(app, application)
	if err != nil {
		println("Some number of operations failed, exiting...")
		return err
	}
	return nil
}

func createClientApplication(userCommands prompt.UserSelections) platform.Application {
	clientApplication := platform.Application{}
	clientApplication.Environments = []string{"dev", "stage", "prod"}
	clientApplication.ApplicationType = userCommands.ClientLanguageType
	clientApplication.ProjectName = userCommands.ProjectName + "-client"
	clientApplication.GitToken = userCommands.GitToken
	clientApplication.PlatformToken = userCommands.PlatformToken
	clientApplication.PlatformTeamName = userCommands.PlatformTeamName
	clientApplication.VersionControl = userCommands.VersionControl
	clientApplication.Platform = userCommands.ApplicationType
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
	serverApplication.VersionControl = userCommands.VersionControl
	serverApplication.Platform = userCommands.ApplicationType
	return serverApplication
}

// NewWebType init constructor
func NewWebType(client applications.App, server applications.App) WebType {
	return WebType{client, server}
}
