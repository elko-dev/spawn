package web

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// Factory to create Web application
type Factory struct {
	serverFactory ServerAppFactory
	clientFactory ClientAppFactory
	webCommand    Prompt
}

// ClientAppFactory factory to create a client Application
type ClientAppFactory interface {
	Create(applicationType string) (applications.Project, error)
}

// ServerAppFactory factory to create a server Application
type ServerAppFactory interface {
	Create(applicationType string) (applications.Project, error)
}

// Prompt interface defines user prompts to determine application configuration
type Prompt interface {
	ForClientType(applicationType string) (string, error)
	ForServerType() (string, error)
}

// Create Web type
func (factory Factory) Create(applicationType string) WebType {

	clientApplicationType, _ := factory.webCommand.ForClientType(applicationType)
	serverApplicationType, _ := factory.webCommand.ForServerType()

	contextLogger := log.WithFields(log.Fields{
		"applicationType":       applicationType,
		"clientApplicationType": clientApplicationType,
		"serverApplicationType": serverApplicationType,
	})

	contextLogger.Debug("Constructing server application...")
	client, _ := factory.serverFactory.Create(clientApplicationType)

	contextLogger.Debug("Constructing client application...")
	server, _ := factory.clientFactory.Create(serverApplicationType)

	return NewWebType(client, server)
}

// NewWebFactory init function
func NewWebFactory(serverFactory ServerAppFactory, clientFactory ClientAppFactory, webCommand Prompt) Factory {
	return Factory{serverFactory, clientFactory, webCommand}
}
