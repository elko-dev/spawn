package gitlab

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/git/local"
)

type GitlabRepo struct {
	HTTP   HTTP
	Git    local.Git
	prompt Prompt
}

type Prompt interface {
	forGitToken() (string, error)
	forGroupId() (string, error)
}

// HTTP describing the functionality to Create repositories
type HTTP interface {
	PostGitRepository(repositoryName string, gitToken string) (GitRepository, error)
	AddEnvironmentVariables(platformToken string, projectID string, gitToken string) error
}

// CreateGitRepository creates gitlab instance
func (git GitlabRepo) CreateGitRepository(repositoryName string, templateURL string, platformToken string) (applications.GitResult, error) {

	gitToken, err := git.prompt.forGitToken()
	repository, err := git.HTTP.PostGitRepository(repositoryName, gitToken)

	if err != nil {
		return applications.GitResult{}, err
	}

	err = git.HTTP.AddEnvironmentVariables(platformToken, repository.ID.String(), gitToken)
	if err != nil {
		println("Failed to add environment variables to Gitlab repo...")
		return applications.GitResult{}, err
	}

	return git.Git.DuplicateRepo(templateURL, gitToken, repository.Name, repository.URL)

}

// NewGitlabRepo Init
func NewGitlabRepo(prompt Prompt) GitlabRepo {
	http := NewGitlabHTTP(prompt)
	return GitlabRepo{http, local.NewLocal(), prompt}
}
