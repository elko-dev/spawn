package api

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// APIType struct to create an application type
type APIType struct {
	Server applications.Project
}

// Create sets up a new API application
func (apiType APIType) Create() error {
	log.WithFields(log.Fields{}).Debug("Creating API")
	return apiType.Server.Create()
}

// NewAPIType init constructor
func NewAPIType(server applications.Project) applications.Project {
	return APIType{server}
}
