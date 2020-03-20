package firebase

import (
	"errors"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/spawnhttp"

	log "github.com/sirupsen/logrus"
)

const (
	webURL = "https://dev-autoenrollfirebase.herokuapp.com/spawn/projects/:id/firebase/web"
)

// WebRequest to create a web app in firebase
type WebRequest struct {
	DisplayName string `json:"displayName"`
}

// WebResponse from firebase
type WebResponse struct {
	AppID       string `json:"appId"`
	ProjectID   string `json:"projectID"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// WebClient to create Web app for firebase
type WebClient struct {
	http http.Client
}

// Create web application in firebase
func (client WebClient) Create(firebaseProjectID string, request WebRequest) (applications.WebApp, error) {
	req, err := createWebRequest(firebaseProjectID, request)
	if err != nil {
		return applications.WebApp{}, err
	}
	resp, err := client.http.Do(req)

	if err != nil {
		return applications.WebApp{}, err
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"response":          resp,
			"firebaseProjectID": firebaseProjectID,
		}).Error("Error creating Web App")

		return applications.WebApp{}, errors.New("Received error creating web project with status ")
	}

	webResponse := WebResponse{}
	err = spawnhttp.MarshalResponse(resp, &webResponse)

	if err != nil {
		return applications.WebApp{}, err
	}

	response := applications.WebApp{
		AppID:       webResponse.AppID,
		Name:        webResponse.Name,
		DisplayName: webResponse.DisplayName,
		ProjectID:   webResponse.ProjectID,
	}
	log.WithFields(log.Fields{
		"rawResponse": resp,
		"webResponse": response,
	}).Debug("Successfully created Web Project")

	return response, nil

}

func createWebRequest(firebaseProjectID string, request WebRequest) (*http.Request, error) {
	requestURL := strings.Replace(webURL, ":id", firebaseProjectID, 1)
	return spawnhttp.CreateRequest(requestURL, request)
}

// NewWebClient init
func NewWebClient() WebClient {
	return WebClient{}
}
