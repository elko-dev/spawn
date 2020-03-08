package gcp

import (
	"errors"
	"net/http"

	"github.com/elko-dev/spawn/spawnhttp"
	log "github.com/sirupsen/logrus"
)

const (
	url = "https://dev-autoenrollfirebase.herokuapp.com/spawn/projects"
)

// ProjectRequest to create project
type ProjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Project asdf
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProjectClient to create GCP Project
type ProjectClient struct {
	http http.Client
}

// Create GCP Project
func (client ProjectClient) Create(project ProjectRequest) (Project, error) {
	req, err := createRequest(project)

	if err != nil {
		return Project{}, err
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return Project{}, err
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"request":        req,
			"response":       resp,
			"projectRequest": project,
		}).Error("Error creating GCP project")

		return Project{}, errors.New("Received error creating gcp project with status ")
	}

	projectResponse := Project{}
	spawnhttp.MarshalResponse(resp, projectResponse)
	if err != nil {
		return Project{}, err
	}
	log.WithFields(log.Fields{
		"rawResponse":    resp,
		"projectRequest": project,
		"projectReponse": projectResponse,
	}).Debug("Successfully created GCP Project")

	return projectResponse, nil
}

func createRequest(project ProjectRequest) (*http.Request, error) {
	return spawnhttp.CreateRequest(url, project)
}

// NewProjectClient init
func NewProjectClient() ProjectClient {
	return ProjectClient{http.Client{}}
}
