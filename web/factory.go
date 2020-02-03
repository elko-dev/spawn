package web

import "github.com/elko-dev/spawn/applications"

// Factory to create Web application
type Factory struct {
	nodeJsFactory AppFactory
	// reactFactory  AppFactory
	webCommand Prompt
}

// AppFactory factory to create an Application
type AppFactory interface {
	Create() (applications.Project, error)
}

// Prompt interface defines user prompts to determine application configuration
type Prompt interface {
	ForClientType(applicationType string) (string, error)
	ForServerType() (string, error)
}

// Create Web type
func (factory Factory) Create(applicationType string) WebType {
	//These are no-ops to present to user until more languages are supported
	factory.webCommand.ForClientType(applicationType)
	factory.webCommand.ForServerType()
	//TODO: refactor this to use react once that is implemented
	client, _ := factory.nodeJsFactory.Create()
	server, _ := factory.nodeJsFactory.Create()
	return NewWebType(client, server)
}

// NewWebFactory init function
func NewWebFactory(nodeJsFactory AppFactory, webCommand Prompt) Factory {
	return Factory{nodeJsFactory, webCommand}
}