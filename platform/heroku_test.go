package platform

import (
	"strconv"
	"strings"
	"testing"
)

func TestHerokuNameReturnsCorrectFormat(t *testing.T) {

	actual := createHerokuName("TESTAPP")
	size := len(createHerokuName("TESTAPP"))

	if size != 7 {
		t.Log("expected size to be 8 but was " + strconv.Itoa(size))
	}
	println("actual " + actual)
	if !strings.HasPrefix(actual, "testapp") {
		t.Log("expected " + actual + " to start with testapp")
	}
}
