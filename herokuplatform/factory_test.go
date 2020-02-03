package herokuplatform

import (
	"reflect"
	"testing"

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

	mockPrompt.EXPECT().forEnvironments().Return(expected, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(herokuTeamName, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(token, nil)

	factory := NewFactory(mockPrompt)

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

	mockPrompt.EXPECT().forEnvironments().Return(envs, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(expected, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(token, nil)

	factory := NewFactory(mockPrompt)

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

	mockPrompt.EXPECT().forEnvironments().Return(envs, nil)
	mockPrompt.EXPECT().forHerokuTeamName().Return(herokuTeamName, nil)
	mockPrompt.EXPECT().forPlatformToken().Return(expected, nil)

	factory := NewFactory(mockPrompt)

	platform, _ := factory.Create(projectName, applicationType)
	heroku := platform.(Heroku)
	actual := heroku.platformToken
	if !reflect.DeepEqual(actual, expected) {
		t.Log("platformTeamName did not match")
		t.Fail()
	}

}
