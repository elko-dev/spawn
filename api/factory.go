package api

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Factory to create Web application
type Factory struct {
	serverFactory web.ServerAppFactory
	webCommand    web.Prompt
}

// Create Web type
func (factory Factory) Create(applicationType string) (applicationtype.ApplicationType, error) {

	serverApplicationType, _ := factory.webCommand.ForServerType()

	contextLogger := log.WithFields(log.Fields{
		"applicationType":       applicationType,
		"serverApplicationType": serverApplicationType,
	})

	contextLogger.Debug("Constructing server application...")
	server, _ := factory.serverFactory.Create(serverApplicationType)

	return NewAPIType(server), nil
}

// NewFactory init function
func NewFactory(serverFactory web.ServerAppFactory, webCommand web.Prompt) applicationtype.APITypeFactory {
	return Factory{serverFactory, webCommand}
}
