package web

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/applicationtype"
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
	IncludeBackend() (bool, error)
}

// Create Web type
func (factory Factory) Create(applicationType string) (applicationtype.ApplicationType, error) {
	includeBackend, _ := factory.webCommand.IncludeBackend()

	var clientApplicationType string
	var serverApplicationType string
	var server applications.Project

	clientApplicationType, _ = factory.webCommand.ForClientType(applicationType)
	client, _ := factory.clientFactory.Create(clientApplicationType)

	if includeBackend {
		serverApplicationType, _ = factory.webCommand.ForServerType()
		server, _ = factory.serverFactory.Create(serverApplicationType)
	}

	contextLogger := log.WithFields(log.Fields{
		"applicationType":       applicationType,
		"clientApplicationType": clientApplicationType,
		"serverApplicationType": serverApplicationType,
	})

	contextLogger.Debug("Constructing application...")

	return NewWebType(client, server, includeBackend), nil
}

// NewWebFactory init function
func NewWebFactory(serverFactory ServerAppFactory, clientFactory ClientAppFactory, webCommand Prompt) applicationtype.WebTypeFactory {
	return Factory{serverFactory, clientFactory, webCommand}
}
