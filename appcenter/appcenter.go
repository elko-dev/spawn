package appcenter

import (
	"context"

	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/organization"
)

// Platform struct to create AppCenter
type Platform struct {
	orgClient        organization.Client
	appClient        apps.Client
	organizationName string
	projectName      string
}

// Create AppCenter config
func (platform Platform) Create() error {
	// create organization
	ctx := context.Background()

	_, err := platform.orgClient.CreateOrganization(ctx, &organization.CreateOrganizationArgs{
		DisplayName: &platform.organizationName,
		Name:        &platform.organizationName,
	})

	if err != nil {
		return err
	}
	// create app
	appDescription := "Mobile application"
	os := "Android"
	platformType := "React-Native"
	releastType := "Production"
	_, err = platform.appClient.CreateApp(ctx, &apps.CreateAppArgs{
		DisplayName: &platform.projectName,
		Name:        &platform.projectName,
		Description: &appDescription,
		OS:          &os,
		Platform:    &platformType,
		ReleaseType: &releastType,
	}, platform.organizationName)
	// create distribution group

	// build app?
	return err
}

// NewPlatform init
func NewPlatform(orgClient organization.Client, appClient apps.Client, organizationName string, projectName string) Platform {
	return Platform{orgClient, appClient, organizationName, projectName}
}
