package applications

import (
	"reflect"
	"testing"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/git"
	"github.com/elko-dev/spawn/platform"
)

func TestCreateAppCreateNodeAppWhenGivenNodeParam(t *testing.T) {
	application := platform.Application{ApplicationType: "NodeJs", VersionControl: constants.Gitlab}
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

func TestCreateAppReturnsADOSGitWhenNodeAndADOSIsProvided(t *testing.T) {
	application := platform.Application{ApplicationType: constants.NodeServerType, VersionControl: constants.ADOS}
	app, err := CreateApp(application)
	gitRepo := app.(NodeJs).Repo
	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isADOSType(gitRepo) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(gitRepo))
		t.Fail()
	}
}

func TestCreateAppReturnsGitlabGitWhenNodeAndGitlabIsProvided(t *testing.T) {
	application := platform.Application{ApplicationType: constants.NodeServerType, VersionControl: constants.Gitlab}
	app, err := CreateApp(application)
	gitRepo := app.(NodeJs).Repo
	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isisGitlabType(gitRepo) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(gitRepo))
		t.Fail()
	}
}

func TestCreateAppReturnsFunctionsWhenNodeAndFunctionsIsProvided(t *testing.T) {
	application := platform.Application{ApplicationType: constants.NodeServerType, VersionControl: constants.Gitlab, Platform: constants.AzureFunctions}
	app, err := CreateApp(application)
	gitRepo := app.(NodeJs).Repo
	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isisGitlabType(gitRepo) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(gitRepo))
		t.Fail()
	}
}

func TestCreateAppContainsHerokuPlatformWhenWebPlatformProvided(t *testing.T) {
	application := platform.Application{ApplicationType: constants.NodeServerType, VersionControl: constants.Gitlab, Platform: constants.WebApplicationType}
	app, err := CreateApp(application)
	platform := app.(NodeJs).Platform
	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isHerokuType(platform) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(platform))
		t.Fail()
	}
}

func TestCreateAppContainsFunctionsPlatformWhenFunctionsPlatformProvided(t *testing.T) {
	application := platform.Application{ApplicationType: constants.NodeServerType, VersionControl: constants.Gitlab, Platform: constants.AzureFunctions}
	app, err := CreateApp(application)
	platform := app.(NodeJs).Platform
	if err != nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if !isFunctionsType(platform) {
		t.Log("Incorrect application type returned got ", reflect.TypeOf(platform))
		t.Fail()
	}
}
func TestCreateAppCreateReactAppWhenGivenReactParam(t *testing.T) {
	application := platform.Application{ApplicationType: "React", VersionControl: constants.Gitlab}
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

func TestCreateAppReturnsErrorWhenInvalidVersionControlProvided(t *testing.T) {
	application := platform.Application{ApplicationType: "React", VersionControl: "not found"}
	_, err := CreateApp(application)

	if err == nil {
		t.Log("error returned when non expected ", err)
		t.Fail()
	}

	if err.Error() != "Invalid Git Repository" {
		t.Log("Incorrect error message returned, got ", err.Error())
		t.Fail()
	}
}
func TestCreateAppReturnsErrorWhenInvalidApplicationTypeProvided(t *testing.T) {
	application := platform.Application{ApplicationType: "Does Not exist", VersionControl: constants.Gitlab}
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

func isFunctionsType(t interface{}) bool {
	switch t.(type) {

	case platform.Functions:
		return true
	default:
		return false
	}

}

func isHerokuType(t interface{}) bool {
	switch t.(type) {

	case platform.HerokuPlatform:
		return true
	default:
		return false
	}

}
func isisGitlabType(t interface{}) bool {
	switch t.(type) {

	case git.GitlabRepository:
		return true
	default:
		return false
	}

}
func isADOSType(t interface{}) bool {
	switch t.(type) {

	case git.ADOSRepository:
		return true
	default:
		return false
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
