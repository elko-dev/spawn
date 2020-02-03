package clip

import (
	"testing"

	"github.com/elko-dev/spawn/applicationtype"
	"github.com/urfave/cli"
)

func ExpectedNameIsSpawn(t *testing.T) {

	expected := cli.App{
		Name: "spawn",
	}

	actual := Init(applicationtype.Factory{})

	if expected.Name != actual.Name {
		t.Errorf("Init: != Name (%v != %v)", actual, expected)
	}

}
