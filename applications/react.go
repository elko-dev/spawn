package applications

import "github.com/elko-dev/spawn/platform"

// React struct to create Node aplication
type React struct {
	Repo     GitRepository
	Platform PlatformRepository
}

const (
	reactTemplateURL = "https://github.com/elko-dev/react-template.git"
)

// Create is a function to generate a react application
func (react React) Create(application platform.Application) error {

	err := createApp(react.Platform, application)
	if err != nil {
		return err
	}
	gitRepo, err := react.Repo.CreateGitRepository(application.ProjectName, application.GitToken, application.PlatformToken, reactTemplateURL)
	if err != nil {
		return err
	}

	println("Created gitlab respository with url: ", gitRepo.URL)
	return nil
}

// NewReact init function
func NewReact(gitRepository GitRepository, platform PlatformRepository) React {
	react := React{Repo: gitRepository, Platform: platform}
	return react
}
