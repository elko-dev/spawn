package applications

// Package defining interfaces to support app creation

// Factory to define interface to create Apps
type Factory interface {
	Create() (Project, error)
}

// Project interface representing interface to create an application project
type Project interface {
	Create() error
}

// GitRepo describing the functionality to Create repositories
type GitRepo interface {
	CreateGitRepository(repositoryName string, url string, platformToken string) error
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create() error
}

// PlatformFactory interface defining platform request
type PlatformFactory interface {
	Create(projectName string, applicationType string) (PlatformRepository, error)
}

// GitFactory to create GitRepo
type GitFactory interface {
	Create(projectName string) (GitRepo, error)
}
