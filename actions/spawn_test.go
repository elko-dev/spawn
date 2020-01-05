package actions

import (
	"errors"
	"testing"

	"gitlab.com/shared-tool-chain/spawn/commands"
	"gitlab.com/shared-tool-chain/spawn/git/api"
)

const (
	gitURL                = "SOMEURL"
	expectedPlatformError = "EXPECTED ERROR"
)

type mockPlatform struct {
}

func (mockPlatform mockPlatform) Create(accessToken string, applicationName string, teamName string) (string, error) {
	if gitURL != gitURL {
		return "", errors.New("INCORRECT URL PASSED TO CREATE")
	}
	return "", errors.New(expectedPlatformError)
}

type MockRepository struct {
}

func (mock MockRepository) CreateGitRepository(repositoryName string, accessToken string) (api.GitRepository, error) {
	return api.GitRepository{}, errors.New("GITLAB_ERROR")
}

type MockGoodRepository struct {
}

func (mock MockGoodRepository) CreateGitRepository(repositoryName string, accessToken string) (api.GitRepository, error) {
	return api.GitRepository{URL: gitURL}, nil
}

func TestApplicationReturnsErrorWhenGitlabReturnsError(t *testing.T) {
	mockRepo := MockGoodRepository{}
	mockPlatform := mockPlatform{}
	spawn := SpawnAction{Repo: mockRepo, Platform: mockPlatform}
	expected := "GITLAB_ERROR"

	actual := spawn.Application(commands.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
	}
}

func TestApplicationReturnsErrorWhenGitlabReturnsSuccessfullyButHerokuFails(t *testing.T) {
	mockRepo := MockGoodRepository{}
	mockPlatform := mockPlatform{}
	spawn := SpawnAction{Repo: mockRepo, Platform: mockPlatform}
	expected := expectedPlatformError

	actual := spawn.Application(commands.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
	}
}
