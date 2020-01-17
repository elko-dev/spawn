package applications

import (
	"errors"

	"github.com/elko-dev/spawn/git"
	"github.com/elko-dev/spawn/git/api"
	"github.com/elko-dev/spawn/herokus"
	"github.com/elko-dev/spawn/platform"
)

// Application is a struct representing a full application
type Application struct {
	ProjectName     string
	DeployToken     string
	AccessToken     string
	PlatformName    string
	ApplicationType string
	Environments    []string
}

// App interface representing interface to create an app
type App interface {
	Create(environments []string) error
}

// GitRepository describing the functionality to Create repositories
type GitRepository interface {
	CreateGitRepository(repositoryName string, accessToken string, deployToken string, url string) (api.GitRepository, error)
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create(application herokus.Application, environments []string) error
}

// CreateApp returns an app
func CreateApp(application Application) (App, error) {
	if application.ApplicationType == "NodeJs" {
		nodeJsApp := NewNodeJs(git.NewGitlabRepository(git.NewLocal()), platform.NewHerokuPlatform(), application)
		return nodeJsApp, nil
	}
	if application.ApplicationType == "React" {
		reactApp := NewReact(git.NewGitlabRepository(git.NewLocal()), platform.NewHerokuPlatform(), application)
		return reactApp, nil
	}
	return nil, errors.New("Invalid Application Type")
}

func createApp(platform PlatformRepository, environments []string, application herokus.Application) error {
	err := platform.Create(application, environments)
	if err != nil {
		return err
	}

	return nil
}
