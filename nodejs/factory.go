package nodejs

import (
	"github.com/elko-dev/spawn/applications"
)

// Prompt for Node specific configuration
type Prompt interface {
	ForAppName() (string, error)
}

// Factory to construct Node App
type Factory struct {
	gitFactory      applications.GitFactory
	platformFactory applications.PlatformFactory
	prompt          Prompt
}

// Create method to construct a Project
func (factory Factory) Create() (applications.Project, error) {
	projectName, err := factory.prompt.ForAppName()
	if err != nil {
		return nil, err
	}
	git, err := factory.gitFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	platform, err := factory.platformFactory.Create(projectName)

	if err != nil {
		return nil, err
	}

	return NewNode(git, platform, projectName), nil
}

// NewFactory init func
func NewFactory(gitFactory applications.GitFactory, platformFactory applications.PlatformFactory, prompt Prompt) Factory {
	return Factory{gitFactory, platformFactory, prompt}
}
