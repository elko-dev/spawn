package web

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// WebType struct to create an application type
type WebType struct {
	Client         applications.Project
	Server         applications.Project
	includeBackend bool
}

// Create sets up a new application
func (webType WebType) Create() error {
	log.WithFields(log.Fields{}).Debug("Creating client app")

	err := webType.Client.Create()

	if err != nil {
		return err
	}
	if !webType.includeBackend {
		return nil
	}
	log.WithFields(log.Fields{}).Debug("Creating backend app")

	return webType.Server.Create()
}

// NewWebType init constructor
func NewWebType(client applications.Project, server applications.Project, includeBackend bool) WebType {
	return WebType{client, server, includeBackend}
}
