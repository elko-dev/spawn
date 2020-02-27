package apps

import (
	"testing"
)

func TestNewClientSubstitutesOrgNameInUrl(t *testing.T) {
	orgName := "testorg"
	url := createOrganizationURL(orgName)
	expected := "orgs/testorg/apps"

	if url != expected {
		t.Log("Incorrect url, got " + url + " expected " + expected)
	}
}
