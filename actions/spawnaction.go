package actions

import "github.com/elko-dev/spawn/applications"

const (
	// NodeGraphQLApplicationType is a nodejs application
	NodeGraphQLApplicationType = "NodeJs"
	// ReactApplicationType is a React application
	ReactApplicationType = "React"
)

// SpawnAction struct to leverage Gitlab
type SpawnAction struct {
}

// Application action to create a project Scaffolding
func (spawn SpawnAction) Application(app applications.App, environments []string) error {
	return app.Create(environments)
}

// NewSpawnAction init function
func NewSpawnAction() SpawnAction {
	return SpawnAction{}
}
