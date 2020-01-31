package nodejs

import "github.com/elko-dev/spawn/applications"

// Node struct to create node Project
type Node struct {
	repo        applications.GitRepository
	platform    applications.PlatformRepository
	projectName string
}

// Create  Node Project
func (node Node) Create() error {
	return nil
}

// NewNode init function
func NewNode(repo applications.GitRepository, platform applications.PlatformRepository, projectName string) Node {
	return Node{repo, platform, projectName}
}
