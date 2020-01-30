package prompt

import (
	"errors"
	"testing"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/mocks"
	"github.com/golang/mock/gomock"
)

const (
	applicationType    = "WEB"
	serverType         = constants.NodeServerType
	clientLanguageType = constants.ReactClientLanguageType
	platformName       = "Heroku"
	platformTeamName   = "teamname"
	platformToken      = "tokenid"
	projectName        = "SomeName"
)

func TestWhenUserSelectsApplicationTypeUserSelectionContainsSaidType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)
	platform.EXPECT().Platform().Return(platformName, "", nil)

	selection := Selection{command, platform}
	expected := applicationType

	application, _ := selection.Application()
	actual := application.ApplicationType

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsApplicationTypeReturnsErrorApplicationErrorsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	selection := Selection{m, platform}

	m.EXPECT().ApplicationType().Return("", errors.New("ERROR ENCOUNTERED"))

	_, error := selection.Application()

	if error == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsServerTypeUserSelectionContainsReturnedServerType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)

	platform.EXPECT().Platform().Return(platformName, "", nil)

	selection := Selection{command, platform}

	expected := serverType

	userSelections, _ := selection.Application()
	actual := userSelections.ServerType

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsServerTypeUserSelectionReturnsErrorApplicationFuncToReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, errors.New("error encountered"))
	selection := Selection{command, platform}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsClientLanguageTypeTypeIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)

	platform.EXPECT().Platform().Return(platformName, "", nil)

	selection := Selection{command, platform}

	expected := clientLanguageType

	userSelections, _ := selection.Application()
	actual := userSelections.ClientLanguageType

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsClientLanguageTypeUserSelectionReturnsErrorApplicationFuncToReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return("", errors.New("error encountered"))
	selection := Selection{command, platform}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsPlatformTokenPlatformTokenIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)
	platform.EXPECT().Platform().Return(platformName, platformToken, nil)

	selection := Selection{command, platform}

	expected := platformName

	userSelections, _ := selection.Application()
	actual := userSelections.PlatformToken

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}
func TestWhenUserSelectsPlatformTeamNameTeamNameIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)
	platform.EXPECT().Platform().Return(platformTeamName, "", nil)

	selection := Selection{command, platform}

	expected := platformTeamName

	userSelections, _ := selection.Application()
	actual := userSelections.PlatformToken

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsPlatformReturnsErrorErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)

	platform.EXPECT().Platform().Return("", "", errors.New("ERROR"))

	selection := Selection{command, platform}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}
func TestWhenUserSelectsProjectNameValueIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	command := mocks.NewMockCommand(ctrl)
	platform := mocks.NewMockPlatformCommand(ctrl)

	command.EXPECT().ApplicationType().Return(applicationType, nil)
	command.EXPECT().ServerType().Return(serverType, nil)
	command.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	command.EXPECT().ProjectName().Return(projectName, nil)

	platform.EXPECT().Platform().Return(platformName, "", nil)

	selection := Selection{command, platform}

	expected := projectName

	userSelections, _ := selection.Application()
	actual := userSelections.ProjectName

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}
