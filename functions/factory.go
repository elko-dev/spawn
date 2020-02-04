package functions

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Factory to create Functions App
type Factory struct {
	serverFactory web.ServerAppFactory
	prompt        Prompt
}

type Prompt interface {
	forFunctionType() (string, error)
}

// Create returns a FunctionType
func (factory Factory) Create(applicationType string) (applicationtype.ApplicationType, error) {
	log.WithFields(log.Fields{"applicationType": applicationType}).Debug("Creating functions server")

	functionType, _ := factory.prompt.forFunctionType()

	if functionType == constants.AzureFunctions {
		log.WithFields(log.Fields{
			"applicationType": applicationType,
			"functionType":    functionType,
		}).Debug("Creating Azure Functions")

	}
	//TODO: Make selection meaningful
	nodeJs, _ := factory.serverFactory.Create(applicationType)
	return NewFunctionsType(nodeJs), nil
}

//NewFactory init function
func NewFactory(prompt Prompt, serverFactory web.ServerAppFactory) applicationtype.FunctionTypeFactory {
	return Factory{serverFactory, prompt}
}
