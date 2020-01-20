package applications

import (
	"github.com/elko-dev/spawn/herokus"
)

const (
	nodeTemplateURL = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
)

// NodeJs struct to create Node aplication
type NodeJs struct {
	Name        string
	AccessToken string
	DeployToken string
	TeamName    string
	TemplateURL string
	Repo        GitRepository
	Platform    PlatformRepository
}

// Create is a function to generate a NodeJS application
func (nodeJs NodeJs) Create(environments []string) error {
	herokuApplication := herokus.Application{Buildpack: "heroku/nodejs", AccessToken: nodeJs.DeployToken, TeamName: nodeJs.TeamName, ApplicationName: nodeJs.Name}

	err := createApp(nodeJs.Platform, environments, herokuApplication)
	if err != nil {
		return err
	}
	gitRepo, err := nodeJs.Repo.CreateGitRepository(nodeJs.Name, nodeJs.AccessToken, nodeJs.DeployToken, nodeTemplateURL)
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
	nodeJs.AccessToken = application.AccessToken
	nodeJs.DeployToken = application.DeployToken
	nodeJs.TeamName = application.PlatformName
	if application.TemplateURL == "" {
		nodeJs.TemplateURL = nodeTemplateURL
	} else {
		nodeJs.TemplateURL = application.TemplateURL
	}
	return nodeJs
}
