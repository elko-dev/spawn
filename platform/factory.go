package platform

import "github.com/elko-dev/spawn/applications"

import "github.com/elko-dev/spawn/constants"

// Factory to create platform
type Factory struct {
	prompt           Prompt
	herokuFactory    HerokuPlatformFactory
	functionsFactory FunctionsPlatformFactory
}

// Prompt for platform details
type Prompt interface {
	forPlatformType() (string, error)
}

// HerokuPlatformFactory builds platforms
type HerokuPlatformFactory interface {
	Create(projectName string, applicationType string) (applications.PlatformRepository, error)
}

// FunctionsPlatformFactory builds platforms
type FunctionsPlatformFactory interface {
	Create(projectName string, applicationType string) (applications.PlatformRepository, error)
}

// Create platform
func (factory Factory) Create(projectName string, applicationType string) (applications.PlatformRepository, error) {
	// select platform
	platformType, err := factory.prompt.forPlatformType()

	if err != nil {
		return nil, err
	}

	if platformType == constants.AzureFunctions {
		return factory.functionsFactory.Create(projectName, applicationType)
	}

	//build platform
	return factory.herokuFactory.Create(projectName, applicationType)
}

// NewFactory init
func NewFactory(prompt Prompt, herokuFactory HerokuPlatformFactory, functionsFactory FunctionsPlatformFactory) applications.PlatformFactory {
	return Factory{prompt, herokuFactory, functionsFactory}
}
