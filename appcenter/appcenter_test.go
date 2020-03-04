package appcenter

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/elko-dev/spawn/appcenter/accounts"
	"github.com/elko-dev/spawn/appcenter/api"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
)

func TestPlatformCreation(t *testing.T) {
	skipFunctional(t)
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	connection := api.NewConnection(os.Getenv("APPCENTER_TOKEN"))
	orgName := "ElkoTestOrganization1"
	orgClient := organization.NewClient(connection)
	appClient := apps.NewClient(connection)
	buildClient := builds.NewClient(connection)
	accountClient := accounts.NewClient(connection)

	members := []string{"andrew.larsen@elko.dev"}
	secret := "secret"
	platform := NewPlatform(orgClient, appClient, buildClient, accountClient, orgName, "testprojectName1", members, secret)

	err := platform.Create("https://github.com/elko-dev/react-native-template.git", "7ba6e41ab3a0f3b3ffc6f65d443f0f02d30ab31f")

	if err != nil {
		t.Log("got error, expected none", err)
		t.Fail()
	}
}

func skipFunctional(t *testing.T) {
	if os.Getenv("APPCENTER_TOKEN") == "" {
		t.Skip("Skipping testing in CI environment")
	}
}
