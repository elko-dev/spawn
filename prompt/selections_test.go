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
	platform           = "Heroku"
	projectName        = "SomeName"
)

func TestWhenUserSelectsApplicationTypeUserSelectionContainsSaidType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return(platform, nil)
	m.EXPECT().ProjectName().Return(projectName, nil)

	selection := Selection{m}
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
	selection := Selection{m}

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

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return(platform, nil)
	m.EXPECT().ProjectName().Return(projectName, nil)

	selection := Selection{m}

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

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, errors.New("error encountered"))
	selection := Selection{m}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsClientLanguageTypeTypeIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return(platform, nil)
	m.EXPECT().ProjectName().Return(projectName, nil)

	selection := Selection{m}

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

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return("", errors.New("error encountered"))
	selection := Selection{m}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsPlatformPlatformIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return(platform, nil)
	m.EXPECT().ProjectName().Return(projectName, nil)

	selection := Selection{m}

	expected := platform

	userSelections, _ := selection.Application()
	actual := userSelections.Platform

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsPlatformReturnsErrorErrorIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return("", errors.New("ERROR"))

	selection := Selection{m}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}
func TestWhenUserSelectsProjectNameValueIsReturned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCommand(ctrl)

	m.EXPECT().ApplicationType().Return(applicationType, nil)
	m.EXPECT().ServerType().Return(serverType, nil)
	m.EXPECT().ClientLanguageType(applicationType).Return(clientLanguageType, nil)
	m.EXPECT().Platform().Return(platform, nil)
	m.EXPECT().ProjectName().Return(projectName, nil)
	selection := Selection{m}

	expected := projectName

	userSelections, _ := selection.Application()
	actual := userSelections.ProjectName

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}
