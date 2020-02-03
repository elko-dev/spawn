package actions

// SpawnAction struct to leverage Gitlab
type SpawnAction struct {
}

// Application action to create a project Scaffolding
func (spawn SpawnAction) Application() error {
	return nil
}

// NewSpawnAction init function
func NewSpawnAction() SpawnAction {
	return SpawnAction{}
}
