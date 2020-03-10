package herokuplatform

import (
	"context"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/elko-dev/spawn/constants"
	heroku "github.com/heroku/heroku-go/v5"
)

const (
	nodeBuildPack  = "heroku/nodejs"
	reactBuildPack = "mars/create-react-app"
	databasePlan   = "heroku-postgresql:hobby-dev"
)

// Heroku Platform
type Heroku struct {
	service          *heroku.Service
	platformToken    string
	environments     []string
	projectName      string
	platformTeamName string
	applicationType  string
}

// Create Heroku Platform
func (h Heroku) Create() error {
	contextLogger := log.WithFields(log.Fields{
		"applicationType":  h.applicationType,
		"environments":     h.environments,
		"projectName":      h.projectName,
		"platformTeamName": h.platformTeamName,
	})
	heroku.DefaultTransport.BearerToken = h.platformToken
	ctx := context.Background()

	region := "us"
	stack := "heroku-18"
	contextLogger.Debug("Creating platform environments")
	for _, environment := range h.environments {
		println("Creating environment for " + environment)

		herokuName := createHerokuName(h.projectName, environment)
		createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &h.platformTeamName}

		app, err := h.service.TeamAppCreate(ctx, createOpts)

		if err != nil {
			contextLogger.Error("TeamAppCreateError ", err)
			return err
		}

		buildPackOps, err := createBuildpack(h.applicationType)

		if err != nil {
			contextLogger.Error("error creating heroku build pack ", err)
			return err
		}
		_, err = h.service.BuildpackInstallationUpdate(ctx, app.ID, buildPackOps)

		if err != nil {
			contextLogger.Error("error configuring build pack")
			return err
		}
		contextLogger.Debug("Created Application for " + environment + " at url " + app.WebURL)
		databaseAddOn := createDatabaseAddOn(herokuName, h.platformTeamName)
		_, err = h.service.AddOnCreate(ctx, app.ID, databaseAddOn)
		if err != nil {
			contextLogger.Error("Error creating addons ", err)
			return err
		}
		contextLogger.Debug("Database created for heroku")

	}
	return nil

}

// GetToken retrieves access token for platform
func (h Heroku) GetToken() string {
	return h.platformToken
}

// GetPlatformType returns type of platform
func (h Heroku) GetPlatformType() string {
	return constants.HerokuPlatform
}

func createDatabaseAddOn(herokuName string, platformTeamName string) heroku.AddOnCreateOpts {
	databaseName := herokuName + "-database"
	herokuAddonDatabase := "DATABASE"

	return heroku.AddOnCreateOpts{
		Name: &databaseName,
		Attachment: &struct {
			Name *string `json:"name,omitempty" url:"name,omitempty,key"`
		}{
			Name: &herokuAddonDatabase,
		},
		Config: map[string]string{
			"db-version": "1.2.3",
		},
		Confirm: &platformTeamName,
		Plan:    databasePlan,
	}
}

func createBuildpack(applicationType string) (heroku.BuildpackInstallationUpdateOpts, error) {
	buildPackName, err := getApplicationBuildpack(applicationType)

	if err != nil {
		println("Invalid application type " + applicationType)
		return heroku.BuildpackInstallationUpdateOpts{}, err
	}
	buildPackOps := heroku.BuildpackInstallationUpdateOpts{
		Updates: []struct {
			Buildpack string `json:"buildpack" url:"buildpack,key"`
		}{}}

	buildPackOps.Updates = append(buildPackOps.Updates, struct {
		Buildpack string `json:"buildpack" url:"buildpack,key"`
	}{
		Buildpack: buildPackName,
	})

	return buildPackOps, nil
}
func getApplicationBuildpack(applicationType string) (string, error) {
	if applicationType == constants.NodeServerType {
		return nodeBuildPack, nil
	}

	if applicationType == constants.ReactClientLanguageType {
		return reactBuildPack, nil
	}

	return "", errors.New("Invalid Application Type")

}

func createHerokuName(applicationName string, environment string) string {
	herokuName := strings.ToLower(environment + "-" + applicationName)
	return herokuName
}

// NewHeroku init function
func NewHeroku(platformToken string, environments []string, projectName string, platformTeamName string, applicationType string) Heroku {
	s := heroku.NewService(heroku.DefaultClient)
	h := Heroku{}
	h.platformToken = platformToken
	h.environments = environments
	h.projectName = projectName
	h.platformTeamName = platformTeamName
	h.applicationType = applicationType
	h.service = s
	return h
}
