package applications

import "github.com/elko-dev/spawn/herokus"

// React struct to create Node aplication
type React struct {
	Name        string
	AccessToken string
	DeployToken string
	TeamName    string
	Repo        GitRepository
	Platform    PlatformRepository
}

// Create is a function to generate a react application
func (react React) Create(environment string) error {
	herokuApplication := herokus.Application{Buildpack: "nodejs", AccessToken: react.AccessToken, TeamName: react.TeamName}
	url, err := react.Platform.Create(herokuApplication)
	if err != nil {
		return err
	}
	println("Created NodeJS Application for " + environment + " with url: " + url)
	gitRepo, err := react.Repo.CreateGitRepository(react.Name, react.AccessToken, react.DeployToken, url)
	if err != nil {
		return err
	}

	println("Created gitlab respository with url: ", gitRepo.URL)
	return nil
}

// NewReact init function
func NewReact(gitRepository GitRepository, platform PlatformRepository, application Application) React {
	react := React{Repo: gitRepository, Platform: platform}
	react.Name = application.ProjectName
	react.AccessToken = application.ProjectName
	react.DeployToken = application.DeployToken
	react.TeamName = application.PlatformName
	return react
}
