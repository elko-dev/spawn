package applications

import (
	"errors"
	"testing"

	"github.com/elko-dev/spawn/git/api"
	"github.com/elko-dev/spawn/platform"
)

const (
	gitURL                = "SOMEURL"
	expectedPlatformError = "EXPECTED ERROR"
)

type mockBadPlatform struct {
}

func (mockBadPlatform mockBadPlatform) Create(application platform.Application) error {
	if gitURL != gitURL {
		return errors.New("INCORRECT URL PASSED TO CREATE")
	}
	return errors.New(expectedPlatformError)
}

type mockGoodPlatform struct {
}

func (mockGoodPlatform mockGoodPlatform) Create(application platform.Application) error {
	return nil
}

type mockBadRepository struct {
}

func (mock mockBadRepository) CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error) {
	return api.GitRepository{}, errors.New("GITLAB_ERROR")
}

type mockGoodRepository struct {
}

func (mock mockGoodRepository) CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error) {
	return api.GitRepository{URL: gitURL}, nil
}

func TestNodeJsCreateReturnsErrorWhenGitlabReturnsError(t *testing.T) {
	mockRepo := mockBadRepository{}
	mockBadPlatform := mockGoodPlatform{}
	nodeJs := NodeJs{Repo: mockRepo, Platform: mockBadPlatform}
	expected := "GITLAB_ERROR"

	actual := nodeJs.Create(platform.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestNodeJsCreateReturnsErrorWhenHerokuFails(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockBadPlatform := mockBadPlatform{}
	nodeJs := NodeJs{Repo: mockRepo, Platform: mockBadPlatform}
	expected := expectedPlatformError
	actual := nodeJs.Create(platform.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestNewNodeJsSetsTemplateUrlWhenNoneProvided(t *testing.T) {
	expected := nodeTemplateURL
	actual := getNodeTemplateURL(nodeTemplateURL)
	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestNewNodeJsSetsTemplateUrlToProvidedUrl(t *testing.T) {
	expected := nodeTemplateURL
	actual := getNodeTemplateURL("")

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}
