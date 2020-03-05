package reactnative

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Prompt for React specific configuration
type Prompt interface {
	forAppName() (string, error)
}

// Factory to construct React App
type Factory struct {
	gitFactory applications.GitFactory
	ciFactory  applications.CIFactory
	prompt     Prompt
}

// Create method to construct a Project
func (factory Factory) Create(applicationType string) (applications.Project, error) {
	log.WithFields(log.Fields{
		"applicationType": applicationType,
	}).Debug("Anout to construct react native application...")

	projectName, err := factory.prompt.forAppName()
	if err != nil {
		return nil, err
	}

	git, err := factory.gitFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	platform, err := factory.ciFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"applicationType": applicationType,
		"projectName":     projectName,
		"git":             git,
	}).Debug("Constructing react native application...")

	return NewReactNative(git, platform, projectName), nil
}

// NewFactory init func
func NewFactory(gitFactory applications.GitFactory, platformFactory applications.CIFactory, prompt Prompt) web.ClientAppFactory {
	return Factory{gitFactory, platformFactory, prompt}
}
