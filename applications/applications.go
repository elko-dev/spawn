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
	RepoID          string
}

// GitRepo describing the functionality to Create repositories
type GitRepo interface {
	CreateGitRepository(repositoryName string, templateURL string, platformToken string, replacements map[string]string) (GitResult, error)
	GetRepoType() string
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
	Create(repoURL string, repoID string, latestGitConfig string, gitType string) error
}

// MobileApps struct containing ios and android app
type MobileApps struct {
	IOS     IOSApp
	Android AndroidApp
}

// IOSApp information
type IOSApp struct {
	ID   string
	Name string
}

// AndroidApp information
type AndroidApp struct {
	ID   string
	Name string
}

// MobilePlatform repository that defines creation of mobile platforms
type MobilePlatform interface {
	Create() (MobileApps, error)
}

// MobilePlatformFactory interface defining mobile platform request
type MobilePlatformFactory interface {
	Create(projectName string, applicationType string) (MobilePlatform, error)
}
