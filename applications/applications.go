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

// GitResult from creating repository
type GitResult struct {
	RepoURL         string
	LatestGitCommit string
}

// GitRepo describing the functionality to Create repositories
type GitRepo interface {
	CreateGitRepository(repositoryName string, templateURL string, platformToken string) (GitResult, error)
}

// PlatformRepository repository that defines creation of Platform repo
type PlatformRepository interface {
	Create() error
	GetToken() string
	GetPlatformType() string
}

// PlatformFactory interface defining platform request
type PlatformFactory interface {
	Create(projectName string, applicationType string) (PlatformRepository, error)
}

// GitFactory to create GitRepo
type GitFactory interface {
	Create(projectName string) (GitRepo, error)
}

// CIFactory to create ci platform
type CIFactory interface {
	Create(projectName string) (CIPlatform, error)
}

// CIPlatform defined CI platform
type CIPlatform interface {
	Create(repoURL string, latestGitConfig string) error
}
