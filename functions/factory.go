package functions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
)

// Factory to create Functions App
type Factory struct {
	// nodeJsFactory web.AppFactory
}

// Create returns a FunctionType
func (factory Factory) Create(projectName string, applicationType string) (applications.PlatformRepository, error) {
	// nodeJs, _ := factory.nodeJsFactory.Create(applicationType)

	return NewFunctionsType(), nil
}

//NewFactory init function
func NewFactory() platform.FunctionsPlatformFactory {
	return Factory{}
}
