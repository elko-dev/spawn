package herokuplatform

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
)

// Factory to create a platform
type Factory struct {
	prompt Prompt
}

// Prompt interface to retrieve values from
type Prompt interface {
	forEnvironments() ([]string, error)
	forHerokuTeamName() (string, error)
	forPlatformToken() (string, error)
}

// Create a platform repo
func (factory Factory) Create(projectName string, applicationType string) (applications.PlatformRepository, error) {
	// fields to create
	// applicationType  string
	envs, err := factory.prompt.forEnvironments()

	if err != nil {
		return Heroku{}, err
	}

	teamName, err := factory.prompt.forHerokuTeamName()
	if err != nil {
		return Heroku{}, err
	}

	token, err := factory.prompt.forPlatformToken()
	if err != nil {
		return Heroku{}, err
	}

	return NewHeroku(token, envs, projectName, teamName, applicationType), nil
}

// NewFactory init
func NewFactory(prompt Prompt) platform.HerokuPlatformFactory {
	return Factory{prompt}
}
