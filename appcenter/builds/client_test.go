package builds

import "testing"

func TestNewClientSubstitutesOrgNameAndProjectNameInUrl(t *testing.T) {
	orgName := "testorg"
	appName := "testapp"
	url := createRepoURL(orgName, appName)
	expected := "apps/testorg/testapp/repo_config"

	if url != expected {
		t.Log("Incorrect url, got " + url + " expected " + expected)
		t.Fail()
	}
}

func TestNewClientSubstitutesOrgNameBranchAndProjectNameInUrl(t *testing.T) {
	orgName := "testorg"
	appName := "testapp"
	branch := "testbranch"
	url := createBuildURL(orgName, appName, branch)
	expected := "apps/testorg/testapp/branches/testbranch/builds"

	if url != expected {
		t.Log("Incorrect url, got " + url + " expected " + expected)
		t.Fail()
	}
}
func TestNewClientSubstitutesOrgNameBranchAndProjectNameInUrlForBuildConfig(t *testing.T) {
	orgName := "testorg"
	appName := "testapp"
	branch := "testbranch"
	url := createConfigBuildURL(orgName, appName, branch)
	expected := "apps/testorg/testapp/branches/testbranch/config"

	if url != expected {
		t.Log("Incorrect url, got " + url + " expected " + expected)
		t.Fail()
	}
}
