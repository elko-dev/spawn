package herokus

// Application struct representing the values of a Heroku Application
type Application struct {
	AccessToken     string
	ApplicationName string
	TeamName        string
	Environment     string
	Buildpack       string
}
