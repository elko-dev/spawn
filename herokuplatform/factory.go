package herokuplatform

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
	log "github.com/sirupsen/logrus"
)

// Factory to create a platform
type Factory struct {
	prompt Prompt
	reader platform.Secrets
}

// Prompt interface to retrieve values from
type Prompt interface {
	forEnvironments() ([]string, error)
	forHerokuTeamName() (string, error)
	forPlatformToken() (string, error)
	forAuthSecretPath() (string, error)
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
	secretPath, err := factory.prompt.forAuthSecretPath()
	log.WithFields(log.Fields{
		"projectName": projectName,
		"secretPath":  secretPath,
	}).Debug("Retrieved secret path")

	authSecretFileString, err := factory.reader.AsBase64String(secretPath)
	if err != nil {
		return nil, err
	}

	credentials := "config/spawn-platform.json"

	return NewHeroku(token, envs, projectName, teamName, applicationType, createConfigVars(credentials, authSecretFileString)), nil
}

func createConfigVars(credentials string, authSecretFileString string) map[string]*string {
	typeORMLogging := "true"
	typeORMSynchronize := "true"
	nodeOptions := "--max_old_space_size=4096"
	configVars := make(map[string]*string)
	configVars["GOOGLE_APPLICATION_CREDENTIALS"] = &credentials
	configVars["AUTH_CONFIG"] = &authSecretFileString
	configVars["TYPEORM_SYNCHRONIZE"] = &typeORMSynchronize
	configVars["TYPEORM_LOGGING"] = &typeORMLogging
	configVars["NODE_OPTIONS"] = &nodeOptions

	return configVars
}

// NewFactory init
func NewFactory(prompt Prompt, reader platform.Secrets) platform.HerokuPlatformFactory {
	return Factory{prompt, reader}
}
