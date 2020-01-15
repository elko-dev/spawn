package clip

import (
	"testing"

	"github.com/urfave/cli"
	"github.com/elko-dev/spawn/actions"
)

func ExpectedNameIsSpawn(t *testing.T) {

	expected := cli.App{
		Name: "spawn",
	}

	actual := Init(actions.SpawnAction{})

	if expected.Name != actual.Name {
		t.Errorf("Init: != Name (%v != %v)", actual, expected)
	}

}
