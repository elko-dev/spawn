package git

import (
	"gitlab.com/shared-tool-chain/spawn/actions"
	"gitlab.com/shared-tool-chain/spawn/git/api"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// GitlabRepository struct to leverage Gitlab
type GitlabRepository struct {
	HTTP HTTP
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

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: accessToken,
		},
	})

	if err != nil {
		println("Clone failed")
		println(err.Error())
		return api.GitRepository{}, err
	}

	err = r.DeleteRemote("origin")
	if err != nil {
		println("Delete failed")
		println(err.Error())
		return api.GitRepository{}, err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{repository.URL},
	})
	if err != nil {
		println("Create remote failed")
		println(err.Error())
		return api.GitRepository{}, err
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: accessToken,
		},
	})

	if err != nil {
		println("Push failed")
		println(err.Error())
		return api.GitRepository{}, err
	}

	return repository, nil
}

// NewGitlabRepository init method
func NewGitlabRepository() actions.GitRepository {
	http := api.GitlabHTTP{}
	return GitlabRepository{HTTP: http}
}
