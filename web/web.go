package web

import (
	"github.com/elko-dev/spawn/applications"
)

// WebType struct to create an application type
type WebType struct {
	Client applications.Project
	Server applications.Project
}

// Create sets up a new application
func (webType WebType) Create() error {
	//create client
	//create server
	return nil
}

// NewWebType init constructor
func NewWebType(client applications.Project, server applications.Project) WebType {
	return WebType{client, server}
}
