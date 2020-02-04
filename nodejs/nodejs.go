package nodejs

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	log "github.com/sirupsen/logrus"
)

const (
	herokuTemplateURL    = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
	functionsTemplateURL = "https://github.com/elko-dev/nodejs-azure-functions-template.git"
)

// Node struct to create node Project
type Node struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
}

// Create  Node Project
func (node Node) Create() error {

	templateURL := getTemplateURL(node.platform.GetPlatformType())

	log.WithFields(log.Fields{
		"projectName": node.projectName,
		"templateURL": templateURL,
	}).Debug("Creating NodeJS Git repository")

	err := node.repo.CreateGitRepository(node.projectName, templateURL, node.platform.GetToken())
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{}).Debug("Creating NodeJS platform")

	return node.platform.Create()
}

func getTemplateURL(platformType string) string {
	if platformType == constants.AzureFunctions {
		return functionsTemplateURL
	}

	return herokuTemplateURL

}

// NewNode init function
func NewNode(repo applications.GitRepo, platform applications.PlatformRepository, projectName string) Node {
	return Node{repo, platform, projectName}
}
