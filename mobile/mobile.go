package mobile

import (
	"github.com/elko-dev/spawn/applications"
	log "github.com/sirupsen/logrus"
)

// MobileType struct to create a mobile application type
type MobileType struct {
	Client applications.Project
	Server applications.Project
}

// Create sets up a new application
func (mobile MobileType) Create() error {
	log.WithFields(log.Fields{}).Debug("Creating client app")

	err := mobile.Client.Create()

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{}).Debug("Creating client app")

	return mobile.Server.Create()
}

// NewMobileType init constructor
func NewMobileType(client applications.Project, server applications.Project) MobileType {
	return MobileType{client, server}
}
