package herokuplatform

import (
	"context"
	"errors"
	"strings"

	"github.com/elko-dev/spawn/constants"
	heroku "github.com/heroku/heroku-go/v5"
)

const (
	nodeBuildPack  = "heroku/nodejs"
	nodeTemplate   = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
	reactTemplate  = "https://github.com/elko-dev/react-template.git"
	reactBuildPack = "mars/create-react-app"
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
	heroku.DefaultTransport.BearerToken = h.platformToken

	region := "us"
	stack := "heroku-18"

	for _, environment := range h.environments {
		herokuName := createHerokuName(h.projectName, environment)
		createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &h.platformTeamName}

		app, err := h.service.TeamAppCreate(context.TODO(), createOpts)

		if err != nil {
			println(err.Error())
			return errors.New("Error Creating App")
		}

		buildPackOps, err := createBuildpack(h.applicationType)

		if err != nil {
			println("error creating heroku build pack")
			return err
		}
		_, err = h.service.BuildpackInstallationUpdate(context.TODO(), app.ID, buildPackOps)

		if err != nil {
			println("error configuring build pack")
			return err
		}
		println("Created Application for " + environment + " at url " + app.WebURL)
	}

	return nil

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