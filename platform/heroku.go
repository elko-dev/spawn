package platform

import (
	"context"
	"errors"
	"strings"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/prompt"
	heroku "github.com/heroku/heroku-go/v5"
)

const (
	nodeBuildPack  = "heroku/nodejs"
	nodeTemplate   = "https://github.com/elko-dev/nodejs-graphql-typescript-template.git"
	reactTemplate  = "https://github.com/elko-dev/react-template.git"
	reactBuildPack = "mars/create-react-app"
)

// Application is a struct representing a full application
type Application struct {
	ProjectName     string
	PlatformToken   string
	GitToken        string
	ApplicationType string
	Environments    []string
	TemplateURL     string
}

// HerokuPlatform struct for heroku operations implementation
type HerokuPlatform struct {
	Service *heroku.Service
}

// HerokuApp struct representing the values of a Heroku Application
type HerokuApp struct {
}

// Create method to create heroku repository
func (h HerokuPlatform) Create(application Application) error {
	heroku.DefaultTransport.BearerToken = application.PlatformToken

	region := "us"
	stack := "heroku-18"
	teamName, err := prompt.HerokuTeamName()

	if err != nil {
		println(err.Error())
		return errors.New("Error Retrieving Heroku Team Name")
	}

	for _, environment := range application.Environments {
		herokuName := createHerokuName(application.ProjectName, environment)
		createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &teamName}

		app, err := h.Service.TeamAppCreate(context.TODO(), createOpts)

		if err != nil {
			println(err.Error())
			return errors.New("Error Creating App")
		}

		buildPackOps, err := createBuildpack(application)

		if err != nil {
			println("error creating heroku build pack")
			return err
		}
		_, err = h.Service.BuildpackInstallationUpdate(context.TODO(), app.ID, buildPackOps)

		if err != nil {
			println("error configuring build pack")
			return err
		}
		println("Created Application for " + environment + " at url " + app.WebURL)
	}

	return nil
}

func createBuildpack(application Application) (heroku.BuildpackInstallationUpdateOpts, error) {
	buildPackName, err := getApplicationBuildpack(application.ApplicationType)

	if err != nil {
		println("Invalid application type " + application.ApplicationType)
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

// NewHerokuPlatform init function
func NewHerokuPlatform() HerokuPlatform {
	s := heroku.NewService(heroku.DefaultClient)
	return HerokuPlatform{Service: s}
}
