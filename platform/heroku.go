package platform

import (
	"context"
	"errors"
	"strings"

	"github.com/elko-dev/spawn/herokus"
	"github.com/elko-dev/spawn/prompt"
	heroku "github.com/heroku/heroku-go/v5"
)

// HerokuAPI describing the functionality to interact with Heroku
type HerokuAPI interface {
	CreateApp(name string, region string, stack string)
}

// HerokuPlatform struct for heroku operations implementation
type HerokuPlatform struct {
	Service *heroku.Service
}

// HerokuApp struct representing the values of a Heroku Application
type HerokuApp struct {
}

// Create method to create heroku repository
func (h HerokuPlatform) Create(application herokus.Application, environments []string) error {
	heroku.DefaultTransport.BearerToken = application.AccessToken

	region := "us"
	stack := "heroku-18"
	teamName, err := prompt.HerokuTeamName()

	if err != nil {
		println(err.Error())
		return errors.New("Error Retrieving Heroku Team Name")
	}

	for _, environment := range environments {
		herokuName := createHerokuName(application.ApplicationName, environment)
		createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &teamName}

		app, err := h.Service.TeamAppCreate(context.TODO(), createOpts)

		if err != nil {
			println(err.Error())
			return errors.New("Error Creating App")
		}

		buildPackOps := heroku.BuildpackInstallationUpdateOpts{
			Updates: []struct {
				Buildpack string `json:"buildpack" url:"buildpack,key"`
			}{}}

		buildPackOps.Updates = append(buildPackOps.Updates, struct {
			Buildpack string `json:"buildpack" url:"buildpack,key"`
		}{
			Buildpack: application.Buildpack,
		})

		_, err = h.Service.BuildpackInstallationUpdate(context.TODO(), app.ID, buildPackOps)

		if err != nil {
			println(err.Error())
			return errors.New("Error")
		}
		println("Created Application for " + environment + " at url " + app.WebURL)
	}

	return nil
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
