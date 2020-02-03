package nodejs

import (
	"github.com/elko-dev/spawn/applications"
)

// Node struct to create node Project
type Node struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
	gitToken    string
}

// Create  Node Project
func (node Node) Create() error {
	return nil
}

// NewNode init function
func NewNode(repo applications.GitRepo, platform applications.PlatformRepository, projectName string, gitToken string) Node {
	return Node{repo, platform, projectName, gitToken}
}
