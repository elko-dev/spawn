package react

import "github.com/elko-dev/spawn/applications"

const (
	templateURL             = "https://github.com/elko-dev/react-template.git"
	templateNameReplacement = "myapp"
)

type React struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
}

// Create  React Project
func (react React) Create() error {
	err := react.platform.Create()
	if err != nil {
		return err
	}
	_, err = react.repo.CreateGitRepository(react.projectName, templateURL, react.platform.GetToken(), createReplacements(react.projectName))
	return err
}

func createReplacements(projectName string) map[string]string {
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName
	return replacements
}

// NewReact init function
func NewReact(repo applications.GitRepo, platform applications.PlatformRepository, projectName string) React {
	return React{repo, platform, projectName}
}
