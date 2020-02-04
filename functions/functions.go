package functions

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// FunctionsType struct to create an Azure Functions type
type FunctionsType struct {
	server applications.Project
}

const (
	templateURL = "https://github.com/elko-dev/nodejs-azure-functions-template.git"
)

// Create sets up a new application
func (function FunctionsType) Create() error {
	log.WithFields(log.Fields{}).Debug("Running server creation")

	return function.server.Create()
}

// GetToken retrieves access token for platform
func (function FunctionsType) GetToken() string {
	return ""
}

// NewFunctionsType init function
func NewFunctionsType(server applications.Project) FunctionsType {
	return FunctionsType{server}
}
