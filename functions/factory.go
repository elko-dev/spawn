package functions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/web"
)

// Factory to create Functions App
type Factory struct {
	nodeJsFactory web.AppFactory
}

// Create returns a FunctionType
func (factory Factory) Create(projectName string, applicationType string) (applications.PlatformRepository, error) {
	nodeJs, _ := factory.nodeJsFactory.Create()

	return NewFunctionsType(nodeJs), nil
}

//NewFactory init function
func NewFactory(nodeJsFactory web.AppFactory) platform.FunctionsPlatformFactory {
	return Factory{nodeJsFactory}
}
