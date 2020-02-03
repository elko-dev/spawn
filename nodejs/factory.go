package nodejs

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/web"
)

// Prompt for Node specific configuration
type Prompt interface {
	forAppName() (string, error)
	forPlatform() (string, error)
	forVersionControl() (string, error)
}

// Factory to construct Node App
type Factory struct {
	gitFactory      applications.GitFactory
	platformFactory applications.PlatformFactory
	prompt          Prompt
}

// Create method to construct a Project
func (factory Factory) Create() (applications.Project, error) {
	projectName, err := factory.prompt.forAppName()
	if err != nil {
		return nil, err
	}
	git, err := factory.gitFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	platform, err := factory.platformFactory.Create(projectName, "TODO: ADD ME ")

	if err != nil {
		return nil, err
	}

	return NewNode(git, platform, projectName, ""), nil
}

// NewFactory init func
func NewFactory(gitFactory applications.GitFactory, platformFactory applications.PlatformFactory, prompt Prompt) web.AppFactory {
	return Factory{gitFactory, platformFactory, prompt}
}
