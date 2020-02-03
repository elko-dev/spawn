package nodejs

import (
	"github.com/elko-dev/spawn/applications"
)

const (
	templateURL = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
)

// Node struct to create node Project
type Node struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
}

// Create  Node Project
func (node Node) Create() error {
	err := node.platform.Create()
	if err != nil {
		return err
	}
	return node.repo.CreateGitRepository(node.projectName, templateURL, node.platform.GetToken())
}

// NewNode init function
func NewNode(repo applications.GitRepo, platform applications.PlatformRepository, projectName string) Node {
	return Node{repo, platform, projectName}
}
