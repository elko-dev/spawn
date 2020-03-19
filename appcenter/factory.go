package appcenter

import (
	"github.com/elko-dev/spawn/appcenter/accounts"
	"github.com/elko-dev/spawn/appcenter/api"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/file"
	log "github.com/sirupsen/logrus"
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
	forExternalUserID() (string, error)
}

// Create appcenter factory
func (factory Factory) Create(projectName string) (applications.CIPlatform, error) {
	log.WithFields(log.Fields{
		"projectName": projectName,
	}).Debug("Creating Appcenter")

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

	gitUserID, err := factory.prompt.forExternalUserID()
	if err != nil {
		return nil, err
	}

	secretPath, err := factory.prompt.forAuthSecretPath()
	log.WithFields(log.Fields{
		"projectName": projectName,
		"secretPath":  secretPath,
	}).Debug("Retrieved secret path")

	authSecretFileString, err := factory.reader.AsBase64String(secretPath)
	if err != nil {
		return nil, err
	}

	connection := api.NewConnection(token)
	orgClient := organization.NewClient(connection)
	appClient := apps.NewClient(connection)
	buildClient := builds.NewClient(connection)
	accountClient := accounts.NewClient(connection)

	platform := NewPlatform(orgClient,
		appClient,
		buildClient,
		accountClient,
		orgName,
		projectName,
		members,
		authSecretFileString,
		gitUserID)

	return platform, nil
}

// NewFactory init
func NewFactory(prompt Prompt) applications.CIFactory {
	return Factory{prompt, file.Reader{}}
}
