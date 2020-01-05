package git

import (
	"gitlab.com/shared-tool-chain/spawn/actions"
	"gitlab.com/shared-tool-chain/spawn/git/api"
)

// GitlabRepository struct to leverage Gitlab
type GitlabRepository struct {
	HTTP HTTP
	Git  Git
}

// Git to interact with git
type Git interface {
	DuplicateRepo(url string, accessToken string, repository api.GitRepository) error
}

// HTTP describing the functionality to Create repositories
type HTTP interface {
	PostGitRepository(repositoryName string, accessToken string) (api.GitRepository, error)
	AddEnvironmentVariables(deployToken string, projectID string, accessToken string) error
}

// CreateGitRepository action to create a Gitlab repo
func (gitlab GitlabRepository) CreateGitRepository(repositoryName string, accessToken string, deployToken string) (api.GitRepository, error) {
	repository, err := gitlab.HTTP.PostGitRepository(repositoryName, accessToken)

	err = gitlab.HTTP.AddEnvironmentVariables(deployToken, repository.ID.String(), accessToken)
	if err != nil {
		println("Failed to add environment variables to Gitlab repo...")
		return api.GitRepository{}, err
	}

	println("Added environment variables to Gitlab repo...")

	url := "https://gitlab.com/shared-tool-chain/react-native-template.git"

	if err != nil {
		return api.GitRepository{}, err
	}

	err = gitlab.Git.DuplicateRepo(url, accessToken, repository)

	if err != nil {
		return api.GitRepository{}, err
	}

	return repository, nil
}

// NewGitlabRepository init method
func NewGitlabRepository(git Git) actions.GitRepository {
	http := api.GitlabHTTP{}
	return GitlabRepository{HTTP: http, Git: git}
}
