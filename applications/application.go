package applications

import (
	"errors"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/git"
	"github.com/elko-dev/spawn/platform"
)

// CreateApp returns an app
func CreateApp(application platform.Application) (App, error) {
	gitRepo, err := getRepositoryType(application)

	if err != nil {
		return nil, err
	}

	platform, err := getPlatformType(application.Platform)

	if err != nil {
		return nil, err
	}

	if application.ApplicationType == constants.NodeServerType {
		nodeJsApp := NewNodeJs(gitRepo, platform)
		return nodeJsApp, nil
	}

	if application.ApplicationType == constants.ReactClientLanguageType {
		reactApp := NewReact(gitRepo, platform)
		return reactApp, nil
	}

	return nil, errors.New("Invalid Application Type")
}

func getPlatformType(platformType string) (PlatformRepository, error) {
	if platformType == constants.AzureFunctions {
		return platform.Functions{}, nil
	}

	return platform.NewHerokuPlatform(), nil
}

func getRepositoryType(application platform.Application) (GitRepository, error) {
	if application.VersionControl == constants.ADOS {
		return git.ADOSRepository{}, nil
	}

	if application.VersionControl == constants.Gitlab {
		return git.NewGitlabRepository(git.NewLocal()), nil
	}

	return nil, errors.New("Invalid Git Repository")
}

func createApp(platform PlatformRepository, application platform.Application) error {
	err := platform.Create(application)
	if err != nil {
		return err
	}

	return nil
}
