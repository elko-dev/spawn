package appcenter

import (
	"testing"

	"github.com/elko-dev/spawn/appcenter/api"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/organization"
)

func TestPlatformCreation(t *testing.T) {
	connection := api.NewConnection("")
	orgName := "ElkoTestOrganization"
	orgClient := organization.NewClient(connection)
	appClient := apps.NewClient(connection, orgName)
	platform := NewPlatform(orgClient, appClient, orgName, "testprojectName")

	err := platform.Create()

	if err != nil {
		t.Log("got error, ecpected none", err)
		t.Fail()
	}
}
