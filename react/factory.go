package react

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/web"
)

// Prompt for React specific configuration
type Prompt interface {
	forAppName() (string, error)
	forWebId() (string, error)
}

// Factory to construct React App
type Factory struct {
	gitFactory      applications.GitFactory
	platformFactory applications.PlatformFactory
	prompt          Prompt
}

// Create method to construct a Project
func (factory Factory) Create(applicationType string) (applications.Project, error) {
	projectName, err := factory.prompt.forAppName()
	if err != nil {
		return nil, err
	}
	git, err := factory.gitFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	platform, err := factory.platformFactory.Create(projectName, applicationType)

	if err != nil {
		return nil, err
	}

	webConfigID, err := factory.prompt.forWebId()

	return NewReact(git, platform, projectName, webConfigID), nil
}

// NewFactory init func
func NewFactory(gitFactory applications.GitFactory, platformFactory applications.PlatformFactory, prompt Prompt) web.ClientAppFactory {
	return Factory{gitFactory, platformFactory, prompt}
}
