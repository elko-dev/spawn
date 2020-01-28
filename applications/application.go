package applications

import (
	"errors"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/git"
	"github.com/elko-dev/spawn/git/api"
	"github.com/elko-dev/spawn/platform"
)

// App interface representing interface to create an app
type App interface {
	Create(application platform.Application) error
}

// GitRepository describing the functionality to Create repositories
type GitRepository interface {
	CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error)
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create(application platform.Application) error
}

// CreateApp returns an app
func CreateApp(application platform.Application) (App, error) {
	if application.ApplicationType == constants.NodeServerType {
		nodeJsApp := NewNodeJs(git.NewGitlabRepository(git.NewLocal()), platform.NewHerokuPlatform())
		return nodeJsApp, nil
	}
	if application.ApplicationType == constants.ReactClientLanguageType {
		reactApp := NewReact(git.NewGitlabRepository(git.NewLocal()), platform.NewHerokuPlatform())
		return reactApp, nil
	}
	return nil, errors.New("Invalid Application Type")
}

func createApp(platform PlatformRepository, application platform.Application) error {
	err := platform.Create(application)
	if err != nil {
		return err
	}

	return nil
}
