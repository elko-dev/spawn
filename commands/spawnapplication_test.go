package commands

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/elko-dev/spawn/applications"
)

type mockSpawnAction struct {
}

// SpawnAction describing the functionality to Create repositories
func (mock mockSpawnAction) Application(app applications.App, environments []string) error {
	return errors.New("RUNTIME_ERROR")
}

func TestRunEncountersErrorProcessExitWithCode1(t *testing.T) {
	spawnAction := mockSpawnAction{}
	application := applications.Application{}
	if os.Getenv("BE_CRASHER") == "1" {
		executeAction(spawnAction, application)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestRunEncountersErrorProcessExitWithCode1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}
