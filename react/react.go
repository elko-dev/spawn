package react

import "github.com/elko-dev/spawn/applications"

const (
	templateURL              = "https://github.com/elko-dev/react-template.git"
	templateWebIDReplacement = "1:1095344010247:web:6e16064718094a3b8a66a2"
	templateNameReplacement  = "myapp"
)

type React struct {
	repo        applications.GitRepo
	platform    applications.PlatformRepository
	projectName string
	webConfigId string
}

// Create  React Project
func (react React) Create() error {
	err := react.platform.Create()
	if err != nil {
		return err
	}
	_, err = react.repo.CreateGitRepository(react.projectName, templateURL, react.platform.GetToken(), createReplacements(react.projectName, react.webConfigId))
	return err
}

func createReplacements(projectName string, webConfigID string) map[string]string {
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName
	replacements[templateWebIDReplacement] = webConfigID
	return replacements
}

// NewReact init function
func NewReact(repo applications.GitRepo, platform applications.PlatformRepository, projectName string, webConfigId string) React {
	return React{repo, platform, projectName, webConfigId}
}
