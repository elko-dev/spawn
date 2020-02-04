package azurefunctions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
)

// Factory struct to create ADOS function
type Factory struct {
	prompt Prompt
}

// Prompt for user details
type Prompt interface {
	forToken() (string, error)
	forOrganization() (string, error)
}

// Create ados function
func (factory Factory) Create(projectName string, applicationType string) (applications.PlatformRepository, error) {
	organization, err := factory.prompt.forOrganization()

	if err != nil {
		return nil, err
	}

	token, err := factory.prompt.forToken()

	if err != nil {
		return nil, err
	}

	return NewAzureFunctions(organization, token, projectName), nil
}

// NewFactory init
func NewFactory(prompt Prompt) platform.FunctionsPlatformFactory {
	return Factory{prompt}
}
