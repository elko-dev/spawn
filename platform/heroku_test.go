package platform

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
	}
	println("actual " + actual)
	if !strings.HasPrefix(actual, "dev-testapp") {
		t.Log("expected " + actual + " to start with testapp")
	}
}
