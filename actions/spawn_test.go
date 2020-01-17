package actions

import (
	"errors"
	"testing"
)

const (
	firstEnvironment       = "dev"
	secondEnvironment      = "stage"
	secondEnvironmentError = "EXPECTED_ERROR"
)

type someApp struct {
}

func (someApp someApp) Create(environments []string) error {
	return errors.New("GITLAB_ERROR")
}
func TestApplicationReturnsErrorWhenNodeJsReturnsError(t *testing.T) {
	someApp := someApp{}
	spawn := SpawnAction{}
	expected := "GITLAB_ERROR"
	environments := []string{firstEnvironment}
	actual := spawn.Application(someApp, environments).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()
	}
}
