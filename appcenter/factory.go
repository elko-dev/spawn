package appcenter

import (
	"github.com/elko-dev/spawn/appcenter/accounts"
	"github.com/elko-dev/spawn/appcenter/api"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/file"
)

// Factory struct to create ADOS function
type Factory struct {
	prompt Prompt
	reader file.Reader
}

// Prompt for user details
type Prompt interface {
	forOrganization() (string, error)
	forToken() (string, error)
	forMembers() ([]string, error)
	forAuthSecretPath() (string, error)
}

// Create appcenter factory
func (factory Factory) Create(projectName string) (applications.CIPlatform, error) {
	token, err := factory.prompt.forToken()
	if err != nil {
		return nil, err
	}

	orgName, err := factory.prompt.forOrganization()

	if err != nil {
		return nil, err
	}

	members, err := factory.prompt.forMembers()

	if err != nil {
		return nil, err
	}

	secretPath, err := factory.prompt.forAuthSecretPath()

	authSecretFileString, err := factory.reader.AsString(secretPath)
	if err != nil {
		return nil, err
	}

	connection := api.NewConnection(token)
	orgClient := organization.NewClient(connection)
	appClient := apps.NewClient(connection)
	buildClient := builds.NewClient(connection)
	accountClient := accounts.NewClient(connection)

	platform := NewPlatform(orgClient, appClient, buildClient, accountClient, orgName, projectName, members, authSecretFileString)

	return platform, nil
}

// NewFactory init
func NewFactory(prompt Prompt) applications.CIFactory {
	return Factory{prompt, file.Reader{}}
}
