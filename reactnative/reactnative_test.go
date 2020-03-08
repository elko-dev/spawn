package reactnative

import (
	"reflect"
	"testing"
)

func Test_createReplacements(t *testing.T) {
	projectName := "somename"
	iosID := "ios"
	androidID := "android"
	expectedReplacements := make(map[string]string)
	expectedReplacements[templateNameReplacement] = projectName
	expectedReplacements[templateAndroidIDReplacement] = androidID
	expectedReplacements[templateIOSIDReplacement] = iosID
	actual := createReplacements(projectName, iosID, androidID)

	if !reflect.DeepEqual(actual, expectedReplacements) {
		t.Log("expected", expectedReplacements, " got ", actual)
		t.Fail()
		return
	}
}
