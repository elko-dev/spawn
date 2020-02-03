package web

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// WebType struct to create an application type
type WebType struct {
	Client applications.Project
	Server applications.Project
}

// Create sets up a new application
func (webType WebType) Create() error {
	log.WithFields(log.Fields{}).Debug("Creating client app")

	err := webType.Client.Create()

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{}).Debug("Creating client app")

	return webType.Server.Create()
}

// NewWebType init constructor
func NewWebType(client applications.Project, server applications.Project) WebType {
	return WebType{client, server}
}
