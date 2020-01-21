package applications

import (
	"reflect"
	"testing"

	"github.com/elko-dev/spawn/platform"
)

func TestCreateAppCreateNodeAppWhenGivenNodeParam(t *testing.T) {
	application := platform.Application{ApplicationType: "NodeJs"}
	nodeJsApp, err := CreateApp(application)

	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isNodeJsType(nodeJsApp) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(nodeJsApp))
		t.Fail()
	}
}

func TestCreateAppCreateReactAppWhenGivenReactParam(t *testing.T) {
	application := platform.Application{ApplicationType: "React"}
	reactApp, err := CreateApp(application)

	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isReactType(reactApp) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(reactApp))
		t.Fail()
	}
}

func TestCreateAppReturnsErrorWhenInvalidApplicationTypeProvided(t *testing.T) {
	application := platform.Application{ApplicationType: "Does Not exist"}
	_, err := CreateApp(application)

	if err == nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if err.Error() != "Invalid Application Type" {
		t.Log("Incorrect error message returned, got ", err.Error())
		t.Fail()
	}
}

func isNodeJsType(t interface{}) bool {
	switch t.(type) {

	case NodeJs:
		return true
	default:
		return false
	}

}

func isReactType(t interface{}) bool {
	switch t.(type) {

	case React:
		return true
	default:
		return false
	}

}
