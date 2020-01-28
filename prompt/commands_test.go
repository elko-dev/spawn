package prompt

import (
	"reflect"
	"testing"

	"github.com/elko-dev/spawn/constants"
)

func TestWhenWebIsSelectedReturnsClientSelectionsOfReact(t *testing.T) {
	expected := []string{constants.ReactClientLanguageType}

	actual := getClientLangaugeSelections("Web")

	if !reflect.DeepEqual(actual, expected) {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}

}

func TestWhenMobileIsSelectedReturnsClientSelectionsOfReactNative(t *testing.T) {
	expected := []string{"React Native"}

	actual := getClientLangaugeSelections("Mobile")

	if !reflect.DeepEqual(actual, expected) {
		t.Log("Incorrect error, expected ", expected, " got ", actual)
		t.Fail()

	}

}
