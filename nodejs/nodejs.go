package nodejs

import (
	"errors"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	log "github.com/sirupsen/logrus"
)

const (
	graphQLHerokuTemplateURL = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
	expressHerokuTemplateURL = "https://github.com/elko-dev/express-typescript-template.git"
	functionsTemplateURL     = "https://github.com/elko-dev/nodejs-azure-functions-template.git"
	templateNameReplacement  = "myapp"
)

// Node struct to create node Project
type Node struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
	framework   string
}

// Create  Node Project
func (node Node) Create() error {

	templateURL, err := getTemplateURL(node.framework)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"projectName": node.projectName,
		"templateURL": templateURL,
	}).Debug("Creating NodeJS Git repository")

	_, err = node.repo.CreateGitRepository(node.projectName, templateURL, node.platform.GetToken(), createReplacements(node.projectName))
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{}).Debug("Creating NodeJS platform")

	return node.platform.Create()
}

func createReplacements(projectName string) map[string]string {
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName
	return replacements
}

func getTemplateURL(platformType string) (string, error) {
	if platformType == constants.AzureFunctions {
		return functionsTemplateURL, nil
	}
	if platformType == constants.ExpressHerokuPlatform {
		return expressHerokuTemplateURL, nil
	}
	if platformType == constants.GraphQLHerokuPlatform {
		return graphQLHerokuTemplateURL, nil
	}
	return "", errors.New("invalid template")
}

// NewNode init function
func NewNode(repo applications.GitRepo, platform applications.PlatformRepository, projectName string, framework string) Node {
	return Node{repo, platform, projectName, framework}
}
