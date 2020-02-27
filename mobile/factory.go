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
	includeBackend, err := factory.webCommand.IncludeBackend()

	client, _ := factory.clientFactory.Create(clientApplicationType)

	if !includeBackend {

		log.WithFields(log.Fields{
			"applicationType":       applicationType,
			"clientApplicationType": clientApplicationType,
		}).Debug("Constructing client application...")
		return NewMobileType(client, nil, includeBackend), nil
	}

	serverApplicationType, _ := factory.webCommand.ForServerType()

	log.WithFields(log.Fields{
		"applicationType":       applicationType,
		"clientApplicationType": clientApplicationType,
		"serverApplicationType": serverApplicationType,
	}).Debug("Constructing server application...")

	server, _ := factory.serverFactory.Create(serverApplicationType)

	if err != nil {
		return nil, nil
	}

	return NewMobileType(client, server, includeBackend), nil
}

// NewFactory init function
func NewFactory(serverFactory web.ServerAppFactory, clientFactory web.ClientAppFactory, webCommand web.Prompt) applicationtype.MobileTypeFactory {
	return Factory{serverFactory, clientFactory, webCommand}
}
