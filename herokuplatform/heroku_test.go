package herokuplatform

import (
	"strconv"
	"strings"
	"testing"
)

func TestHerokuNameReturnsCorrectFormat(t *testing.T) {

	actual := createHerokuName("TESTAPP", "dev")
	size := len(createHerokuName("TESTAPP", "dev"))

	if size != 11 {
		t.Log("expected size to be 11 but was " + strconv.Itoa(size))
		t.Fail()
	}
	println("actual " + actual)
	if !strings.HasPrefix(actual, "dev-testapp") {
		t.Log("expected " + actual + " to start with testapp")
		t.Fail()
	}
}

func TestCreateBuildpackReturnsCorrectBuildpackValueWhenNodeJsProvided(t *testing.T) {
	buildPack, _ := createBuildpack("NodeJs")
	expected := "heroku/nodejs"
	actual := buildPack.Updates[0].Buildpack
	if actual != expected {
		t.Log("got wrong buildpack, expected " + expected + " got " + actual)
		t.Fail()
	}
}

func TestCreateBuildpackReturnsCorrectBuildpackValueWhenReactProvided(t *testing.T) {
	buildPack, _ := createBuildpack("React")

	actual := buildPack.Updates[0].Buildpack

	if actual != "mars/create-react-app" {
		t.Log("got wrong buildpack, got", actual)
		t.Fail()
	}
}
