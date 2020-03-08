package firebase

import (
	"errors"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/spawnhttp"
)

const (
	androidURL = "https://dev-autoenrollfirebase.herokuapp.com/spawn/projects/:id/firebase/android"
)

// AndroidRequest to create ios app in firebase
type AndroidRequest struct {
	BundleID    string `json:"bundleID"`
	DisplayName string `json:"displayName"`
}

// AndroidResponse from android
type AndroidResponse struct {
	ID       string `json:"id"`
	BundleID string `json:"bundleID"`
	Name     string `json:"name"`
}

// AndroidClient to create android firebase app
type AndroidClient struct {
	http http.Client
}

// Create Android app in firebase
func (client AndroidClient) Create(firebaseProjectID string, request AndroidRequest) (applications.AndroidApp, error) {
	req, err := createAndroidRequest(firebaseProjectID, request)
	if err != nil {
		return applications.AndroidApp{}, nil
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return applications.AndroidApp{}, nil
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"response":          resp,
			"firebaseProjectID": firebaseProjectID,
		}).Error("Error creating Android app")

		return applications.AndroidApp{}, errors.New("Received error creating android project with status ")
	}

	androidResponse := AndroidResponse{}
	err = spawnhttp.MarshalResponse(resp, &androidResponse)

	if err != nil {
		return applications.AndroidApp{}, err
	}
	androidApp := applications.AndroidApp{
		ID:   androidResponse.ID,
		Name: androidResponse.Name,
	}
	log.WithFields(log.Fields{
		"rawResponse":     resp,
		"androidResponse": androidApp,
	}).Debug("Successfully created Android Project")

	return androidApp, nil
}

func createAndroidRequest(firebaseProjectID string, request AndroidRequest) (*http.Request, error) {
	requestURL := strings.Replace(androidURL, ":id", firebaseProjectID, 1)
	return spawnhttp.CreateRequest(requestURL, request)
}

// NewAndroidClient init
func NewAndroidClient() AndroidClient {
	return AndroidClient{}
}
