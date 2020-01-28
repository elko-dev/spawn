package actions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
)

// SpawnAction struct to leverage Gitlab
type SpawnAction struct {
}

// Application action to create a project Scaffolding
func (spawn SpawnAction) Application(app applications.App, application platform.Application) error {
	return app.Create(application)
}

// NewSpawnAction init function
func NewSpawnAction() SpawnAction {
	return SpawnAction{}
}
