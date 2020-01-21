package applications

import "github.com/elko-dev/spawn/platform"

// React struct to create Node aplication
type React struct {
	Name        string
	AccessToken string
	DeployToken string
	Repo        GitRepository
	Platform    PlatformRepository
}

const (
	reactTemplateURL = "https://github.com/elko-dev/react-native-template.git"
)

// Create is a function to generate a react application
func (react React) Create(application platform.Application, environments []string) error {

	err := createApp(react.Platform, environments, application)
	if err != nil {
		return err
	}
	gitRepo, err := react.Repo.CreateGitRepository(react.Name, react.AccessToken, react.DeployToken, reactTemplateURL)
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
