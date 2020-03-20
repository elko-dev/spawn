package reactnative

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/web"
	log "github.com/sirupsen/logrus"
)

// Prompt for React specific configuration
type Prompt interface {
	forAppName() (string, error)
	forIncludingPlatform() (bool, error)
}

// Factory to construct React App
type Factory struct {
	gitFactory      applications.GitFactory
	ciFactory       applications.CIFactory
	platformFactory applications.MobilePlatformFactory
	prompt          Prompt
}

// Create method to construct a Project
func (factory Factory) Create(applicationType string) (applications.Project, error) {
	log.WithFields(log.Fields{
		"applicationType": applicationType,
	}).Debug("About to construct react native application...")

	projectName, err := factory.prompt.forAppName()
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"applicationType": applicationType,
		"projectName":     projectName,
	}).Debug("Selected project name ", projectName)

	git, err := factory.gitFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	ciPlatform, err := factory.ciFactory.Create(projectName)
	if err != nil {
		return nil, err
	}

	includePlatform, err := factory.prompt.forIncludingPlatform()
	if err != nil {
		log.WithFields(log.Fields{
			"applicationType": applicationType,
			"projectName":     projectName,
		}).Error("Error selecting Platform ", err)
		return nil, err
	}
	log.WithFields(log.Fields{
		"applicationType": applicationType,
		"projectName":     projectName,
		"includePlatform": includePlatform,
	}).Debug("Including Platform ", includePlatform)

	var platform applications.MobilePlatform

	if includePlatform {
		platform, err = factory.platformFactory.Create(projectName, applicationType)
		if err != nil {
			return nil, err
		}
	}

	log.WithFields(log.Fields{
		"applicationType": applicationType,
		"projectName":     projectName,
		"git":             git,
	}).Debug("Constructing react native application...")

	return NewReactNative(git, ciPlatform, platform, projectName, includePlatform), nil
}

// NewFactory init func
func NewFactory(gitFactory applications.GitFactory,
	ciPlatformFactory applications.CIFactory,
	platformFactory applications.MobilePlatformFactory,
	prompt Prompt) web.ClientAppFactory {
	return Factory{gitFactory, ciPlatformFactory, platformFactory, prompt}
}
