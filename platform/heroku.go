package platform

import (
	"context"
	"errors"
	"strings"

	heroku "github.com/heroku/heroku-go/v5"
	"gitlab.com/shared-tool-chain/spawn/actions"
)

// HerokuAPI describing the functionality to interact with Heroku
type HerokuAPI interface {
	CreateApp(name string, region string, stack string)
}

// HerokuPlatform struct for heroku operations implementation
type HerokuPlatform struct {
	Service *heroku.Service
}

// Create method to create heroku repository
func (h HerokuPlatform) Create(accessToken string, applicationName string, teamName string) (string, error) {
	heroku.DefaultTransport.BearerToken = accessToken

	region := "us"
	stack := "heroku-18"

	herokuName := createHerokuName(applicationName)
	createOpts := heroku.TeamAppCreateOpts{Name: &herokuName, Region: &region, Stack: &stack, Team: &teamName}

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
		Buildpack: "mars/create-react-app",
	})

	_, err = h.Service.BuildpackInstallationUpdate(context.TODO(), app.ID, buildPackOps)

	if err != nil {
		println(err.Error())
		return "", errors.New("Error")
	}

	return app.WebURL, nil
}

func createHerokuName(applicationName string) string {
	// herokuName := applicationName + strconv.Itoa(rand.Intn(3000000))
	herokuName := strings.ToLower(applicationName)
	return herokuName
}

// NewHerokuPlatform init function
func NewHerokuPlatform() actions.PlatformRepository {
	s := heroku.NewService(heroku.DefaultClient)
	return HerokuPlatform{Service: s}
}
