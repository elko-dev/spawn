package git

import (
	"github.com/elko-dev/spawn/git/api"
)

// GitlabRepository struct to leverage Gitlab
type GitlabRepository struct {
	HTTP HTTP
	Git  Git
}

// Git to interact with git
type Git interface {
	DuplicateRepo(url string, gitToken string, repository api.GitRepository) error
}

// HTTP describing the functionality to Create repositories
type HTTP interface {
	PostGitRepository(repositoryName string, gitToken string) (api.GitRepository, error)
	AddEnvironmentVariables(platformToken string, projectID string, gitToken string) error
}

// CreateGitRepository action to create a Gitlab repo
func (gitlab GitlabRepository) CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error) {
	repository, err := gitlab.HTTP.PostGitRepository(repositoryName, gitToken)

	if err != nil {
		return api.GitRepository{}, err
	}

	err = gitlab.HTTP.AddEnvironmentVariables(platformToken, repository.ID.String(), gitToken)
	if err != nil {
		println("Failed to add environment variables to Gitlab repo...")
		return api.GitRepository{}, err
	}

	err = gitlab.Git.DuplicateRepo(url, gitToken, repository)

	if err != nil {
		return api.GitRepository{}, err
	}

	return repository, nil
}

// NewGitlabRepository init method
func NewGitlabRepository(git Git) GitlabRepository {
	http := api.GitlabHTTP{}
	return GitlabRepository{HTTP: http, Git: git}
}
