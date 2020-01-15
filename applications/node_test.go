package applications

import (
	"errors"
	"testing"

	"github.com/elko-dev/spawn/git/api"
	"github.com/elko-dev/spawn/herokus"
)

const (
	gitURL                = "SOMEURL"
	expectedPlatformError = "EXPECTED ERROR"
)

type mockBadPlatform struct {
}

func (mockBadPlatform mockBadPlatform) Create(application herokus.Application) (string, error) {
	if gitURL != gitURL {
		return "", errors.New("INCORRECT URL PASSED TO CREATE")
	}
	return "", errors.New(expectedPlatformError)
}

type mockGoodPlatform struct {
}

func (mockGoodPlatform mockGoodPlatform) Create(application herokus.Application) (string, error) {
	return "URL", nil
}

type mockBadRepository struct {
}

func (mock mockBadRepository) CreateGitRepository(repositoryName string, accessToken string, deployToken string, url string) (api.GitRepository, error) {
	return api.GitRepository{}, errors.New("GITLAB_ERROR")
}

type mockGoodRepository struct {
}

func (mock mockGoodRepository) CreateGitRepository(repositoryName string, accessToken string, deployToken string, url string) (api.GitRepository, error) {
	return api.GitRepository{URL: gitURL}, nil
}

func TestNodeJsCreateReturnsErrorWhenGitlabReturnsError(t *testing.T) {
	mockRepo := mockBadRepository{}
	mockBadPlatform := mockGoodPlatform{}
	nodeJs := NodeJs{Repo: mockRepo, Platform: mockBadPlatform, Name: "", TeamName: "", AccessToken: "", DeployToken: ""}
	expected := "GITLAB_ERROR"
	environments := []string{"dev"}

	actual := nodeJs.Create(environments[0]).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestNodeJsCreateReturnsErrorWhenGitlabReturnsSuccessfullyButHerokuFails(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockBadPlatform := mockBadPlatform{}
	nodeJs := NodeJs{Repo: mockRepo, Platform: mockBadPlatform, Name: "", TeamName: "", AccessToken: "", DeployToken: ""}
	expected := expectedPlatformError
	environments := []string{"dev"}

	actual := nodeJs.Create(environments[0]).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}
