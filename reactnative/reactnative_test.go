package reactnative

import (
	"reflect"
	"testing"
)

func Test_createReplacements_include_mobile(t *testing.T) {
	projectName := "somename"
	iosID := "ios"
	androidID := "android"
	webID := "web"
	expectedReplacements := make(map[string]string)
	expectedReplacements[templateNameReplacement] = projectName
	expectedReplacements[templateAndroidIDReplacement] = androidID
	expectedReplacements[templateIOSIDReplacement] = iosID
	expectedReplacements[templateWebIDReplacement] = webID
	actual := createReplacements(projectName, iosID, androidID, webID, true)

	if !reflect.DeepEqual(actual, expectedReplacements) {
		t.Log("expected", expectedReplacements, " got ", actual)
		t.Fail()
		return
	}
}

func Test_createReplacements_dont_include_mobile(t *testing.T) {
	projectName := "somename"
	iosID := "ios"
	androidID := "android"
	webID := "web"
	expectedReplacements := make(map[string]string)
	expectedReplacements[templateNameReplacement] = projectName
	actual := createReplacements(projectName, iosID, androidID, webID, false)

	if !reflect.DeepEqual(actual, expectedReplacements) {
		t.Log("expected", expectedReplacements, " got ", actual)
		t.Fail()
		return
	}
}
