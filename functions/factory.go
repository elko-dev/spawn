package functions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Factory to create Functions App
type Factory struct {
	serverFactory web.ServerAppFactory
}

// Create returns a FunctionType
func (factory Factory) Create(applicationType string) (applications.PlatformRepository, error) {
	log.WithFields(log.Fields{"applicationType": applicationType}).Debug("Creating functions server")
	nodeJs, _ := factory.serverFactory.Create(applicationType)
	return NewFunctionsType(nodeJs), nil
}

//NewFactory init function
func NewFactory(serverFactory web.ServerAppFactory) platform.FunctionsPlatformFactory {
	return Factory{serverFactory}
}
