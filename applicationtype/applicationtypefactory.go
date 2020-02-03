package applicationtype

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/platform"
	web "github.com/elko-dev/spawn/web"
)

// Prompt interface defines user prompts to determine application type
type Prompt interface {
	ForType() (string, error)
}

//TODO: Setting this here to make progress.  Need to refactor to move this elsewhere and actually define
type ApplicationType interface {
	Create() error
}

// Factory to create an application type
type Factory struct {
	prompt          Prompt
	webFactory      web.Factory
	functionFactory platform.FunctionsPlatformFactory
}

// CreateApplicationType creates app type
func (factory Factory) CreateApplicationType() ApplicationType {
	// prompt user for application type
	appType, _ := factory.prompt.ForType()

	var applicationType ApplicationType

	if appType == constants.WebApplicationType {
		applicationType = factory.webFactory.Create(appType)
	}

	//TODO: COME BACK HERE!!!
	if appType == constants.AzureFunctions {
		applicationType, _ = factory.functionFactory.Create("PROJECT_NAME_CHANGE_ME", appType)
	}

	return applicationType
}

// NewFactory creates an ApplicationType factory
func NewFactory(prompt Prompts, webFactory web.Factory, functionsFactory platform.FunctionsPlatformFactory) Factory {
	return Factory{prompt, webFactory, functionsFactory}
}
