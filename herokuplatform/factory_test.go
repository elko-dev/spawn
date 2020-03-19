package herokuplatform

import (
	"reflect"
	"testing"

	"github.com/elko-dev/spawn/platform"
	gomock "github.com/golang/mock/gomock"
)

const (
	projectName     = "test"
	applicationType = "test"
	herokuTeamName  = "teeamName"
	token           = "token"
)

func TestWhenEnvironmentsAreReturnedHerokuContainsEnvironments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := []string{"dev", "stage"}
	mockPrompt := NewMockPrompt(ctrl)
	mockReader := platform.NewMockSecrets(ctrl)

	mockPrompt.EXPECT().forEnvironments().Return(expected, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(herokuTeamName, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(token, nil)
	mockPrompt.EXPECT().forAuthSecretPath().Return("path", nil)
	mockReader.EXPECT().AsBase64String("path").Return("Secret", nil)
	factory := NewFactory(mockPrompt, mockReader)

	platform, _ := factory.Create(projectName, applicationType)
	heroku := platform.(Heroku)

	actual := heroku.environments
	if !reflect.DeepEqual(actual, expected) {
		t.Log("environments did not match")
		t.Fail()
	}

}
func TestWhenHerokuTeamNameAreSelectedHerokuContainsTeamName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := herokuTeamName
	envs := []string{"dev", "stage"}

	mockPrompt := NewMockPrompt(ctrl)
	mockReader := platform.NewMockSecrets(ctrl)

	mockPrompt.EXPECT().forEnvironments().Return(envs, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(expected, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(token, nil)
	mockPrompt.EXPECT().forAuthSecretPath().Return("path", nil)
	mockReader.EXPECT().AsBase64String("path").Return("Secret", nil)
	factory := NewFactory(mockPrompt, mockReader)

	platform, _ := factory.Create(projectName, applicationType)
	heroku := platform.(Heroku)
	actual := heroku.platformTeamName
	if !reflect.DeepEqual(actual, expected) {
		t.Log("platformTeamName did not match")
		t.Fail()
	}

}

func TestWhenTokenIsSelectedHerokuContainsToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := token
	envs := []string{"dev", "stage"}

	mockPrompt := NewMockPrompt(ctrl)
	mockReader := platform.NewMockSecrets(ctrl)

	mockPrompt.EXPECT().forEnvironments().Return(envs, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(herokuTeamName, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(expected, nil)
	mockPrompt.EXPECT().forAuthSecretPath().Return("path", nil)
	mockReader.EXPECT().AsBase64String("path").Return("Secret", nil)
	factory := NewFactory(mockPrompt, mockReader)

	platform, _ := factory.Create(projectName, applicationType)
	heroku := platform.(Heroku)
	actual := heroku.platformToken
	if !reflect.DeepEqual(actual, expected) {
		t.Log("platformTeamName did not match")
		t.Fail()
	}

}
