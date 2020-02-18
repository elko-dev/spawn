package appcenter

import (
	"context"

	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
)

// Platform struct to create AppCenter
type Platform struct {
	orgClient        organization.Client
	appClient        apps.Client
	buildClient      builds.Client
	organizationName string
	projectName      string
	repoURL          string
	latestGitConfig  string
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
	// TODO: create a team and add app to team
	// create app
	err = createAndroidApp(ctx, &platform)
	if err != nil {
		return err
	}

	return createIOSApp(ctx, &platform)
}

// CreateApp for app center
func (platform Platform) CreateApp(ctx context.Context, description *string, os *string, platformType *string, releaseType *string) error {
	projectName := normalizeProjectName(platform.projectName, *os)
	_, err := platform.appClient.CreateApp(ctx, &apps.CreateAppArgs{
		DisplayName: &projectName,
		Name:        &projectName,
		Description: description,
		OS:          os,
		Platform:    platformType,
		ReleaseType: releaseType,
	}, platform.organizationName)

	if err != nil {
		return err
	}
	_, err = platform.buildClient.ConfigureRepo(ctx, &builds.RepoConfigArgs{
		RepoURL: platform.repoURL,
	}, platform.organizationName, projectName)

	if err != nil {
		return err
	}
	_, err = platform.buildClient.ConfigureBuild(ctx, builds.CreateConfigArgs(), platform.organizationName, projectName)

	if err != nil {
		return err
	}

	_, err = platform.buildClient.Build(ctx, &builds.BuildArgs{
		SourceVersion: platform.latestGitConfig,
		Debug:         true,
	}, platform.organizationName, projectName)

	return err
}

func createAndroidApp(ctx context.Context, platform *Platform) error {
	description := "Mobile application"
	os := "Android"
	platformType := "React-Native"
	releastType := "Production"
	return platform.CreateApp(ctx, &description, &os, &platformType, &releastType)
}

func createIOSApp(ctx context.Context, platform *Platform) error {
	description := "Mobile application"
	os := "iOS"
	platformType := "React-Native"
	releastType := "Production"
	return platform.CreateApp(ctx, &description, &os, &platformType, &releastType)
}

func normalizeProjectName(projectName string, os string) string {
	return projectName + "-" + os
}

// NewPlatform init
func NewPlatform(orgClient organization.Client,
	appClient apps.Client,
	buildClient builds.Client,
	organizationName string,
	projectName string,
	repoURL string,
	latestGitConfig string) Platform {
	return Platform{orgClient, appClient, buildClient, organizationName, projectName, repoURL, latestGitConfig}
}
