package mobile

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Factory to create Web application
type Factory struct {
	serverFactory web.ServerAppFactory
	clientFactory web.ClientAppFactory
	webCommand    web.Prompt
}

// Create Mobile type
func (factory Factory) Create(applicationType string) (applicationtype.ApplicationType, error) {

	clientApplicationType, _ := factory.webCommand.ForClientType(applicationType)
	serverApplicationType, _ := factory.webCommand.ForServerType()

	contextLogger := log.WithFields(log.Fields{
		"applicationType":       applicationType,
		"clientApplicationType": clientApplicationType,
		"serverApplicationType": serverApplicationType,
	})

	contextLogger.Debug("Constructing server application...")
	client, _ := factory.serverFactory.Create(serverApplicationType)

	contextLogger.Debug("Constructing client application...")
	server, _ := factory.clientFactory.Create(clientApplicationType)

	return NewMobileType(client, server), nil
}

// NewFactory init function
func NewFactory(serverFactory web.ServerAppFactory, clientFactory web.ClientAppFactory, webCommand web.Prompt) applicationtype.MobileTypeFactory {
	return Factory{serverFactory, clientFactory, webCommand}
}
