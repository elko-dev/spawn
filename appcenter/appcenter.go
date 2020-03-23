package appcenter

import (
	"context"

	"github.com/elko-dev/spawn/appcenter/accounts"
	"github.com/elko-dev/spawn/appcenter/apps"
	"github.com/elko-dev/spawn/appcenter/builds"
	"github.com/elko-dev/spawn/appcenter/organization"
	"github.com/elko-dev/spawn/constants"

	log "github.com/sirupsen/logrus"
)

const (
	androidKeyStoreKey = "ANDROID_KEYSTORE_KEY"
	authSecretName     = "AUTH_CONFIG"
)

// Platform struct to create AppCenter
type Platform struct {
	orgClient           organization.Client
	appClient           apps.Client
	buildClient         builds.Client
	accountsClient      accounts.Client
	organizationName    string
	projectName         string
	distributionMembers []string
	authSecret          string
	externalUserID      string
}

// Create AppCenter config
func (platform Platform) Create(repoURL string, repoID string, latestGitConfig string, gitType string) error {

	// create organization
	ctx := context.Background()
	log.WithFields(log.Fields{
		"organizationName": platform.organizationName,
		"projectName":      platform.projectName,
		"repoURL":          repoURL,
		"latestGitConfig":  latestGitConfig,
	}).Debug("Creating appcenter organization")

	_, err := platform.orgClient.CreateOrganization(ctx, &organization.CreateOrganizationArgs{
		DisplayName: &platform.organizationName,
		Name:        &platform.organizationName,
	})

	if err != nil {
		return err
	}

	distributionResponse, err := platform.accountsClient.CreateDistributionGroup(ctx, &accounts.DistributionGroupArg{
		DisplayName: platform.organizationName,
		Name:        platform.organizationName,
	}, &platform.organizationName)

	if err != nil {
		return err
	}

	// TODO: create a team and add app to team
	// create app
	androidName, err := createAndroidApp(ctx, &platform, &repoURL, &latestGitConfig, &distributionResponse.ID, repoID, gitType)
	if err != nil {
		return err
	}

	iosName, err := createIOSApp(ctx, &platform, &repoURL, &latestGitConfig, &distributionResponse.ID, repoID, gitType)

	if err != nil {
		return err
	}

	apps := make([]accounts.Apps, 0)
	apps = append(apps, accounts.Apps{Name: androidName})
	apps = append(apps, accounts.Apps{Name: iosName})

	err = platform.accountsClient.CreateAppsDistributionGroup(ctx, &accounts.AppsForDistributionArg{
		Apps: apps,
	}, &platform.organizationName, &platform.organizationName)

	if err != nil {
		return err
	}

	return platform.accountsClient.AddMemberToDistribution(ctx, &accounts.AddMemberArgs{
		UserEmails: platform.distributionMembers,
	}, &platform.organizationName, &platform.organizationName)

}

// CreateApp for app center
func (platform Platform) CreateApp(ctx context.Context,
	description *string,
	os *string,
	platformType *string,
	releaseType *string,
	repoURL *string,
	latestGitConfig *string,
	distributionID *string,
	environmentVariables []builds.EnvironmentVariables,
	repoID string,
	gitType string) (string, error) {

	println("Creating Firebase Project and Apps")
	projectName := normalizeProjectName(platform.projectName, *os)

	_, err := platform.appClient.CreateApp(ctx, &apps.CreateAppArgs{
		DisplayName: &projectName,
		Name:        &projectName,
		Description: description,
		OS:          os,
		Platform:    platformType,
		ReleaseType: releaseType,
	}, platform.organizationName)

	logContext := log.WithFields(log.Fields{
		"organizationName": platform.organizationName,
		"projectName":      platform.projectName,
		"repoURL":          *repoURL,
		"latestGitConfig":  *latestGitConfig,
		"description":      *description,
		"os":               *os,
		"platformType":     *platformType,
	})

	if err != nil {
		logContext.Info("Error creating appcenter app")
		return "", err
	}
	repoConfig := builds.RepoConfigArgs{
		RepoURL: *repoURL,
		RepoID:  repoID,
	}
	if gitType == constants.Gitlab {
		repoConfig.ExternalUserID = platform.externalUserID
	}
	_, err = platform.buildClient.ConfigureRepo(ctx, &repoConfig, platform.organizationName, projectName)

	if err != nil {
		logContext.Info("Error configuring appcenter app")
		return "", err
	}

	args := builds.CreateConfigArgs(distributionID, environmentVariables, &builds.Keystore{
		KeyAlias:         "app",
		KeyPassword:      "abcdef12",
		KeystoreFilename: "my.keystore",
		KeystorePassword: "abcdef12",
	})
	_, err = platform.buildClient.ConfigureBuild(ctx,
		args,
		platform.organizationName,
		projectName)

	if err != nil {
		logContext.Info("Error creating appcenter build")
		return "", err
	}

	_, err = platform.buildClient.Build(ctx, &builds.BuildArgs{
		SourceVersion: *latestGitConfig,
		Debug:         true,
	}, platform.organizationName, projectName)

	return projectName, err
}

func createAndroidApp(ctx context.Context,
	platform *Platform,
	repoURL *string,
	latestGitConfig *string,
	distributionID *string,
	repoID string,
	gitType string) (string, error) {
	description := "Mobile application"
	os := "Android"
	platformType := "React-Native"
	releastType := "Production"
	//TODO: implement me
	encryptToken := "REPLACE ME"
	environmentVariables := []builds.EnvironmentVariables{
		builds.EnvironmentVariables{
			Name:  androidKeyStoreKey,
			Value: encryptToken,
		},
		builds.EnvironmentVariables{
			Name:  authSecretName,
			Value: platform.authSecret,
		}}
	return platform.CreateApp(ctx, &description, &os, &platformType, &releastType, repoURL, latestGitConfig, distributionID, environmentVariables, repoID, gitType)
}

func createIOSApp(ctx context.Context,
	platform *Platform,
	repoURL *string,
	latestGitConfig *string,
	distributionID *string,
	repoID string,
	gitType string) (string, error) {
	description := "Mobile application"
	os := "iOS"
	platformType := "React-Native"
	releastType := "Production"
	environmentVariables := []builds.EnvironmentVariables{
		builds.EnvironmentVariables{
			Name:  authSecretName,
			Value: platform.authSecret,
		},
	}
	return platform.CreateApp(ctx, &description, &os, &platformType, &releastType, repoURL, latestGitConfig, distributionID, environmentVariables, repoID, gitType)
}

func normalizeProjectName(projectName string, os string) string {
	return projectName + "-" + os
}

// NewPlatform init
func NewPlatform(orgClient organization.Client,
	appClient apps.Client,
	buildClient builds.Client,
	accountsClient accounts.Client,
	organizationName string,
	projectName string,
	members []string,
	authSecret string,
	externalUserID string) Platform {
	return Platform{
		orgClient,
		appClient,
		buildClient,
		accountsClient,
		organizationName,
		projectName,
		members,
		authSecret,
		externalUserID,
	}
}
