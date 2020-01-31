package applicationtype

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/functions"
	prompt "github.com/elko-dev/spawn/prompt"
	web "github.com/elko-dev/spawn/web"
)

// Prompt interface defines user prompts to determine application type
type Prompt interface {
	ForType() (string, error)
}

//TODO: Setting this here to make progress.  Need to refactor to move this elsewhere and actually define
type TempAppType interface {
	Create(action web.SpawnAction, userCommands prompt.UserSelections) error
}

// Factory to create an application type
type Factory struct {
	prompt Prompt
}

// CreateApplicationType creates app type
func (factory Factory) CreateApplicationType() TempAppType {
	webFactory := web.Factory{}
	functionFactory := functions.Factory{}
	// prompt user for application type
	appType, _ := factory.prompt.ForType()

	var applicationType TempAppType

	if appType == constants.WebApplicationType {
		applicationType = webFactory.Create(appType)
	}

	if appType == constants.AzureFunctions {
		applicationType = functionFactory.Create()
	}

	return applicationType
}

// NewFactory creates an ApplicationType factory
func NewFactory() Factory {
	return Factory{prompt: Prompts{}}
}
