package applicationtype

import (
	"errors"

	"github.com/elko-dev/spawn/constants"
)

// Prompt interface defines user prompts to determine application type
type Prompt interface {
	ForType() (string, error)
}

//TODO: Setting this here to make progress.  Need to refactor to move this elsewhere and actually define
type ApplicationType interface {
	Create() error
}

//TODO: Setting this here to make progress.  Need to refactor to move this elsewhere and actually define
type WebTypeFactory interface {
	Create(applicationType string) (ApplicationType, error)
}
type MobileTypeFactory interface {
	Create(applicationType string) (ApplicationType, error)
}
type FunctionTypeFactory interface {
	Create(applicationType string) (ApplicationType, error)
}
type APITypeFactory interface {
	Create(applicationType string) (ApplicationType, error)
}

// Factory to create an application type
type Factory struct {
	prompt            Prompt
	webFactory        WebTypeFactory
	functionFactory   FunctionTypeFactory
	mobileTypeFactory MobileTypeFactory
	apiTypeFactory    APITypeFactory
}

// CreateApplicationType creates app type
func (factory Factory) CreateApplicationType() (ApplicationType, error) {
	// prompt user for application type
	appType, _ := factory.prompt.ForType()

	if appType == constants.WebApplicationType {
		applicationType, err := factory.webFactory.Create(appType)
		if err != nil {
			return nil, err
		}
		return applicationType, nil

	}

	if appType == constants.AzureFunctions {
		applicationType, err := factory.functionFactory.Create(appType)
		if err != nil {
			return nil, err
		}
		return applicationType, nil

	}

	if appType == constants.MobileApplicationType {
		applicationType, err := factory.mobileTypeFactory.Create(appType)
		if err != nil {
			return nil, err
		}
		return applicationType, nil
	}
	if appType == constants.APIApplicationType {
		applicationType, err := factory.apiTypeFactory.Create(appType)
		if err != nil {
			return nil, err
		}
		return applicationType, nil
	}

	return nil, errors.New("Invalid selection")
}

// NewFactory creates an ApplicationType factory
func NewFactory(prompt Prompts, webFactory WebTypeFactory, functionsFactory FunctionTypeFactory, mobileTypeFactory MobileTypeFactory, apiTypeFactory APITypeFactory) Factory {
	return Factory{prompt, webFactory, functionsFactory, mobileTypeFactory, apiTypeFactory}
}
