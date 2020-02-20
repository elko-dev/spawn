package reactnative

import "github.com/elko-dev/spawn/applications"

const (
	templateURL = "https://github.com/elko-dev/react-native-template.git"
)

type ReactNative struct {
	repo        applications.GitRepo
	ciPlatform  applications.CIPlatform
	projectName string
}

// Create ReactNative Project
func (react ReactNative) Create() error {
	// TODO: this fails the interface segregration principle.  Need to refactor
	response, err := react.repo.CreateGitRepository(react.projectName, templateURL, "")

	if err != nil {
		return err
	}
	return react.ciPlatform.Create(response.RepoURL, response.LatestGitCommit)
}

// NewReactNative init function
func NewReactNative(repo applications.GitRepo, platform applications.CIPlatform, projectName string) ReactNative {
	return ReactNative{repo, platform, projectName}
}
