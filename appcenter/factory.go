package appcenter

import (
	"github.com/elko-dev/spawn/appcenter/api"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
	"github.com/elko-dev/spawn/applications"
)

// Factory struct to create ADOS function
type Factory struct {
	prompt Prompt
}

// Prompt for user details
type Prompt interface {
	forOrganization() (string, error)
	forToken() (string, error)
}

// Create appcenter factory
func (factory Factory) Create(projectName string) (applications.CIPlatform, error) {
	token, err := factory.prompt.forToken()
	if err != nil {
		return nil, err
	}

	orgName, err := factory.prompt.forOrganization()

	connection := api.NewConnection(token)
	orgClient := organization.NewClient(connection)
	appClient := apps.NewClient(connection)
	buildClient := builds.NewClient(connection)

	platform := NewPlatform(orgClient, appClient, buildClient, orgName, projectName)

	return platform, nil
}

// NewFactory init
func NewFactory(prompt Prompt) applications.CIFactory {
	return Factory{prompt}
}
