package firebase

import (
	"errors"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/spawnhttp"
	log "github.com/sirupsen/logrus"
)

const (
	url = "https://dev-autoenrollfirebase.herokuapp.com/spawn/projects/:id/firebase"
)

type FirebaseProjectResponse struct {
	ID   string
	Name string
}

// ProjectClient to create project
type ProjectClient struct {
	http http.Client
}

type EmptyRequest struct {
}

// Create new Firebase project for a GCP Project
func (client ProjectClient) Create(gcpProjectID string) (FirebaseProjectResponse, error) {
	log.WithFields(log.Fields{
		"gcpProjectID": gcpProjectID,
	}).Debug("Creating Firebase Project")

	req, err := createRequest(gcpProjectID)
	if err != nil {
		return FirebaseProjectResponse{}, err
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return FirebaseProjectResponse{}, err
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"rawRquest":    req,
			"rawResponse":  resp,
			"gcpProjectID": gcpProjectID,
		}).Error("Error creating Firebase project")

		return FirebaseProjectResponse{}, errors.New("Received error creating firebase project with status ")
	}

	return FirebaseProjectResponse{gcpProjectID, gcpProjectID}, nil
}

func createRequest(gcpProjectID string) (*http.Request, error) {
	requestURL := strings.Replace(url, ":id", gcpProjectID, 1)
	return spawnhttp.CreateRequest(requestURL, EmptyRequest{})
}

// NewProjectClient init
func NewProjectClient() FirebaseProject {
	return ProjectClient{}
}
