package applications

import (
	"testing"

	"github.com/elko-dev/spawn/platform"
)

type mockHerokuPlatform struct {
}

func (mockHerokuPlatform mockHerokuPlatform) Create(application platform.Application) error {
	return nil
}

func TestReactCreateReturnsErrorWhenGitlabReturnsError(t *testing.T) {
	mockRepo := mockBadRepository{}
	mockPlatform := mockGoodPlatform{}
	react := React{Repo: mockRepo, Platform: mockPlatform}
	expected := "GITLAB_ERROR"

	actual := react.Create(platform.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestReactCreateReturnsErrorWhenGitlabReturnsSuccessfullyButHerokuFails(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockBadPlatform := mockBadPlatform{}
	react := React{Repo: mockRepo, Platform: mockBadPlatform}
	expected := expectedPlatformError

	actual := react.Create(platform.Application{}).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

// TODO: test to be removed once we factor our platform details
func TestHerokuIsProvidedCorrectBuildPack(t *testing.T) {
	mockRepo := mockGoodRepository{}
	mockHerokuPlatform := mockHerokuPlatform{}
	react := React{Repo: mockRepo, Platform: mockHerokuPlatform}

	error := react.Create(platform.Application{})

	if error != nil {
		t.Log("no error expected")
		t.Fail()

	}
}
