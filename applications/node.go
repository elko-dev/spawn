package applications

import "github.com/elko-dev/spawn/platform"

const (
	nodeTemplateURL = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
)

// NodeJs struct to create Node aplication
type NodeJs struct {
	Repo     GitRepository
	Platform PlatformRepository
}

// Create is a function to generate a NodeJS application
func (nodeJs NodeJs) Create(application platform.Application) error {

	err := createApp(nodeJs.Platform, application)
	if err != nil {
		return err
	}

	templateURL := getNodeTemplateURL(application.TemplateURL)
	gitRepo, err := nodeJs.Repo.CreateGitRepository(application.ProjectName, application.GitToken, application.PlatformToken, templateURL)
	if err != nil {
		return err
	}

	println("Created gitlab respository with url: ", gitRepo.URL)
	return nil
}

func getNodeTemplateURL(templateURL string) string {
	if templateURL == "" {
		return nodeTemplateURL
	}

	return templateURL
}

// NewNodeJs init function
func NewNodeJs(gitRepository GitRepository, platform PlatformRepository) NodeJs {
	nodeJs := NodeJs{Repo: gitRepository, Platform: platform}
	return nodeJs
}
