package applications

import (
	"github.com/elko-dev/spawn/git/api"
	"github.com/elko-dev/spawn/platform"
)

// Package defining interfaces to support app creation

// Factory to define interface to create Apps
type Factory interface {
	Create() (Project, error)
}

// Project interface representing interface to create an application project
type Project interface {
	Create() error
}

// App interface representing interface to create an app
type App interface {
	Create(application platform.Application) error
}

// GitRepository describing the functionality to Create repositories
type GitRepository interface {
	CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error)
}

// GitRepo describing the functionality to Create repositories
type GitRepo interface {
	CreateGitRepository() (api.GitRepository, error)
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create(application platform.Application) error
}

// Platform repository that defines creation of Platform repo
type Platform interface {
	Create() error
}

// PlatformFactory interface defining platform request
type PlatformFactory interface {
	Create(projectName string) (PlatformRepository, error)
}

// GitFactory to create GitRepository
type GitFactory interface {
	Create(projectName string) (GitRepository, error)
}
