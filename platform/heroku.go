package platform

import (
	"context"
	"errors"
	"strings"

	"github.com/elko-dev/spawn/herokus"
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
	AccessToken     string
	ApplicationName string
	TeamName        string
	Environment     string
	Buildpack       string
}

// Create method to create heroku repository
func (h HerokuPlatform) Create(application herokus.Application) (string, error) {
	heroku.DefaultTransport.BearerToken = application.AccessToken

	region := "us"
	stack := "heroku-18"

	herokuName := createHerokuName(application.ApplicationName, application.Environment)
	createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &application.TeamName}

	app, err := h.Service.TeamAppCreate(context.TODO(), createOpts)

	if err != nil {
		println(err.Error())
		return "", errors.New("Error")
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
		return "", errors.New("Error")
	}

	return app.WebURL, nil
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
