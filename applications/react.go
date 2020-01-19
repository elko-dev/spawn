package applications

import "github.com/elko-dev/spawn/herokus"

// React struct to create Node aplication
type React struct {
	Name        string
	AccessToken string
	DeployToken string
	TeamName    string
	Repo        GitRepository
	Platform    PlatformRepository
}

const (
	reactTemplateURL = "https://github.com/elko-dev/react-native-template.git"
)

// Create is a function to generate a react application
func (react React) Create(environments []string) error {
	herokuApplication := herokus.Application{Buildpack: "mars/create-react-app", AccessToken: react.DeployToken, TeamName: react.TeamName, ApplicationName: react.Name}

	err := createApp(react.Platform, environments, herokuApplication)
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
func NewReact(gitRepository GitRepository, platform PlatformRepository, application Application) React {
	react := React{Repo: gitRepository, Platform: platform}
	react.Name = application.ProjectName
	react.AccessToken = application.AccessToken
	react.DeployToken = application.DeployToken
	react.TeamName = application.PlatformName
	return react
}
