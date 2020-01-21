package actions

import (
	"errors"
	"github.com/elko-dev/spawn/platform"
	"testing"
)

const (
	firstEnvironment       = "dev"
	secondEnvironment      = "stage"
	secondEnvironmentError = "EXPECTED_ERROR"
)

type someApp struct {
}

func (someApp someApp) Create(application platform.Application, environments []string) error {
	return errors.New("GITLAB_ERROR")
}
func TestApplicationReturnsErrorWhenNodeJsReturnsError(t *testing.T) {
	someApp := someApp{}
	spawn := SpawnAction{}
	expected := "GITLAB_ERROR"
	environments := []string{firstEnvironment}
	application := platform.Application{}
	actual := spawn.Application(someApp, application, environments).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()
	}
}
