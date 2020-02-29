package directory

import (
	"testing"
)

func TestTemplateDirectory_replace(t *testing.T) {
	template := TemplateDirectory{"test"}
	err := template.replace(".git/config", nil, nil)

	if err != nil {
		t.Log("Got error, none expected ", err)
		t.Fail()
	}
}
