package prompt

import (
	"errors"
	"testing"
)

const (
	applicationType    = "WEB"
	serverType         = "NodeJs"
	clientLanguageType = "React"
)

type goodMockCommand struct {
}

func (goodMockCommand goodMockCommand) ApplicationType() (string, error) {
	return applicationType, nil
}

func (goodMockCommand goodMockCommand) ServerType() (string, error) {
	return serverType, nil
}

func (goodMockCommand goodMockCommand) ClientLanguageType(inputApplicationType string) (string, error) {
	if inputApplicationType != applicationType {
		println("Incorrect application type passed, expected " + applicationType + " got " + inputApplicationType)
		return "", errors.New("Incorrect application type passed, expected " + applicationType + " got " + inputApplicationType)
	}
	return clientLanguageType, nil
}

type badMockCommand struct {
}

func (badMockCommand badMockCommand) ApplicationType() (string, error) {
	return "", errors.New("ERROR ENCOUNTERED")
}

func (badMockCommand badMockCommand) ServerType() (string, error) {
	return "", errors.New("ERROR ENCOUNTERED")
}
func (badMockCommand badMockCommand) ClientLanguageType(applicationType string) (string, error) {
	return serverType, nil
}

type badServerTypeMockCommand struct {
}

func (badServerTypeMockCommand badServerTypeMockCommand) ApplicationType() (string, error) {
	return applicationType, nil
}

func (badServerTypeMockCommand badServerTypeMockCommand) ServerType() (string, error) {
	return "", errors.New("ERROR ENCOUNTERED")
}

func (badServerTypeMockCommand badServerTypeMockCommand) ClientLanguageType(applicationType string) (string, error) {
	return serverType, nil
}

type badClientLanguageMock struct {
}

func (badClientLanguageMock badClientLanguageMock) ApplicationType() (string, error) {
	return applicationType, nil
}

func (badClientLanguageMock badClientLanguageMock) ServerType() (string, error) {
	return serverType, nil
}

func (badClientLanguageMock badClientLanguageMock) ClientLanguageType(applicationType string) (string, error) {
	return "", errors.New("ERROR ENCOUNTERED")
}

func TestWhenUserSelectsApplicationTypeUserSelectionContainsSaidType(t *testing.T) {
	selection := Selection{goodMockCommand{}}
	expected := applicationType

	application, _ := selection.Application()
	actual := application.ApplicationType

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsApplicationTypeReturnsErrorApplicationErrorsError(t *testing.T) {
	selection := Selection{badMockCommand{}}
	_, error := selection.Application()

	if error == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsServerTypeUserSelectionContainsReturnedServerType(t *testing.T) {
	selection := Selection{goodMockCommand{}}
	expected := serverType

	userSelections, _ := selection.Application()
	actual := userSelections.ServerType

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsServerTypeUserSelectionReturnsErrorApplicationFuncToReturnError(t *testing.T) {
	selection := Selection{badServerTypeMockCommand{}}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}

func TestWhenUserSelectsClientLanguageTypeTypeIsReturned(t *testing.T) {
	selection := Selection{goodMockCommand{}}
	expected := clientLanguageType

	userSelections, _ := selection.Application()
	actual := userSelections.ClientLanguageType

	if actual != expected {
		t.Log("Incorrect type, expected ", expected, " got ", actual)
		t.Fail()

	}
}

func TestWhenUserSelectsClientLanguageTypeUserSelectionReturnsErrorApplicationFuncToReturnError(t *testing.T) {
	selection := Selection{badClientLanguageMock{}}

	_, err := selection.Application()

	if err == nil {
		t.Log("Expected error, got none")
		t.Fail()

	}
}
