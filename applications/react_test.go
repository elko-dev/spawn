package applications

import (
	"errors"
	"testing"

	"github.com/elko-dev/spawn/herokus"
)

type mockHerokuPlatform struct {
}

func (mockHerokuPlatform mockHerokuPlatform) Create(application herokus.Application) (string, error) {
	if application.Buildpack != "mars/create-react-app" {
		return "", errors.New("Invalid buildpack")
	}
	return "URL", nil
}

func TestReactCreateReturnsErrorWhenGitlabReturnsError(t *testing.T) {
	mockRepo := mockBadRepository{}
	mockPlatform := mockGoodPlatform{}
	react := React{Repo: mockRepo, Platform: mockPlatform, Name: "", TeamName: "", AccessToken: "", DeployToken: ""}
	expected := "GITLAB_ERROR"
	environments := []string{"dev"}

	actual := react.Create(environments[0]).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestReactCreateReturnsErrorWhenGitlabReturnsSuccessfullyButHerokuFails(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockBadPlatform := mockBadPlatform{}
	react := React{Repo: mockRepo, Platform: mockBadPlatform, Name: "", TeamName: "", AccessToken: "", DeployToken: ""}
	expected := expectedPlatformError
	environments := []string{"dev"}

	actual := react.Create(environments[0]).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

// TODO: test to be removed once we factor our platform details
func TestHerokuIsProvidedCorrectBuildPack(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockHerokuPlatform := mockHerokuPlatform{}
	react := React{Repo: mockRepo, Platform: mockHerokuPlatform, Name: "", TeamName: "", AccessToken: "", DeployToken: ""}
	environments := []string{"dev"}

	error := react.Create(environments[0])

	if error != nil {
		t.Log("no error expected")
		t.Fail()

	}
}
