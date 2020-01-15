package actions

import (
	"errors"
	"testing"
)

type mockNodeJs struct {
}

func (nodeJs mockNodeJs) Create(environment string) error {
	return errors.New("GITLAB_ERROR")
}
func TestApplicationReturnsErrorWhenNodeJsReturnsError(t *testing.T) {
	mockNodeJs := mockNodeJs{}
	spawn := SpawnAction{}
	expected := "GITLAB_ERROR"
	environments := []string{"dev"}
	actual := spawn.Application(mockNodeJs, environments).Error()

	if actual != expected {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()
	}
}
