package gitlab

import "github.com/elko-dev/spawn/git/local"

type GitlabRepo struct {
	HTTP   HTTP
	Git    local.Git
	prompt Prompt
}

type Prompt interface {
	forGitToken() (string, error)
}

// HTTP describing the functionality to Create repositories
type HTTP interface {
	PostGitRepository(repositoryName string, gitToken string) (GitRepository, error)
	AddEnvironmentVariables(platformToken string, projectID string, gitToken string) error
}

// CreateGitRepository creates gitlab instance
func (git GitlabRepo) CreateGitRepository(repositoryName string, url string, platformToken string) error {

	gitToken, err := git.prompt.forGitToken()
	repository, err := git.HTTP.PostGitRepository(repositoryName, gitToken)

	if err != nil {
		return err
	}

	err = git.HTTP.AddEnvironmentVariables(platformToken, repository.ID.String(), gitToken)
	if err != nil {
		println("Failed to add environment variables to Gitlab repo...")
		return err
	}

	err = git.Git.DuplicateRepo(url, gitToken, repository.Name, repository.URL)

	if err != nil {
		return err
	}

	return nil

}

// NewGitlabRepo Init
func NewGitlabRepo(prompt Prompt) GitlabRepo {
	http := GitlabHTTP{}

	return GitlabRepo{http, local.NewLocal(), prompt}
}
