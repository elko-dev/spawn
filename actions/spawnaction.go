package actions

import (
	"gitlab.com/shared-tool-chain/spawn/commands"
	"gitlab.com/shared-tool-chain/spawn/git/api"
)

// GitRepository describing the functionality to Create repositories
type GitRepository interface {
	CreateGitRepository(repositoryName string, accessToken string) (api.GitRepository, error)
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create(accessToken string, applicationName string) (string, error)
}

// SpawnAction struct to leverage Gitlab
type SpawnAction struct {
	Repo     GitRepository
	Platform PlatformRepository
}

// Application action to create a project Scaffolding
func (spawn SpawnAction) Application(application commands.Application) error {
	gitRepo, err := spawn.Repo.CreateGitRepository(application.ProjectName, application.AccessToken)
	if err != nil {
		return err
	}
	println("Created gitlab respository with url: ", gitRepo.URL)
	println("Creating heroku pipeline...")

	url, err := spawn.Platform.Create(application.DeployToken, application.ProjectName)

	if err != nil {
		return err
	}

	println("Created heroku platform with url: ", url)
	return nil
}

// NewSpawnAction init function
func NewSpawnAction(gitRepository GitRepository, platform PlatformRepository) SpawnAction {
	return SpawnAction{gitRepository, platform}
}
