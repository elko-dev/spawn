package applications

import (
	"github.com/elko-dev/spawn/herokus"
)

const (
	url = "https://github.com/elko-dev/react-native-template.git"
)

// NodeJs struct to create Node aplication
type NodeJs struct {
	Name        string
	AccessToken string
	DeployToken string
	TeamName    string
	Repo        GitRepository
	Platform    PlatformRepository
}

// Create is a function to generate a NodeJS application
func (nodeJs NodeJs) Create(environment string) error {
	herokuApplication := herokus.Application{Buildpack: "mars/create-react-app", AccessToken: nodeJs.AccessToken, TeamName: nodeJs.TeamName}
	url, err := nodeJs.Platform.Create(herokuApplication)
	if err != nil {
		return err
	}
	println("Created NodeJS Application for " + environment + " with url: " + url)
	gitRepo, err := nodeJs.Repo.CreateGitRepository(nodeJs.Name, nodeJs.AccessToken, nodeJs.DeployToken, url)
	if err != nil {
		return err
	}

	println("Created gitlab respository with url: ", gitRepo.URL)
	return nil
}

// NewNodeJs init function
func NewNodeJs(gitRepository GitRepository, platform PlatformRepository, application Application) NodeJs {
	nodeJs := NodeJs{Repo: gitRepository, Platform: platform}
	nodeJs.Name = application.ProjectName
	nodeJs.AccessToken = application.ProjectName
	nodeJs.DeployToken = application.DeployToken
	nodeJs.TeamName = application.PlatformName
	return nodeJs
}
