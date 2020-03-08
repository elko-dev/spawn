package firebase

import (
	"errors"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/spawnhttp"
	log "github.com/sirupsen/logrus"
)

const (
	iosURL = "https://dev-autoenrollfirebase.herokuapp.com/spawn/projects/:id/firebase/ios"
)

// IOSRequest to create ios app in firebase
type IOSRequest struct {
	BundleID    string `json:"bundleID"`
	DisplayName string `json:"displayName"`
}

// IOSResponse from firebase
type IOSResponse struct {
	ID       string `json:"id"`
	BundleID string `json:"bundleID"`
	Name     string `json:"name"`
}

// IOSClient to create IOS app for firebase
type IOSClient struct {
	http http.Client
}

// Create ios app in Firebase
func (client IOSClient) Create(firebaseProjectID string, request IOSRequest) (IOSResponse, error) {

	req, err := createIOSRequest(firebaseProjectID, request)
	if err != nil {
		return IOSResponse{}, err
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return IOSResponse{}, err
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"response":          resp,
			"firebaseProjectID": firebaseProjectID,
		}).Error("Error creating IOS App")

		return IOSResponse{}, errors.New("Received error creating ios project with status ")
	}

	iosResponse := IOSResponse{}
	err = spawnhttp.MarshalResponse(resp, iosResponse)

	if err != nil {
		return IOSResponse{}, err
	}

	log.WithFields(log.Fields{
		"rawResponse": resp,
		"iosResponse": resp,
	}).Debug("Successfully created IOS Project")

	return iosResponse, nil
}

func createIOSRequest(firebaseProjectID string, request IOSRequest) (*http.Request, error) {
	requestURL := strings.Replace(iosURL, ":id", firebaseProjectID, 1)
	return spawnhttp.CreateRequest(requestURL, request)
}

// NewIOSClient init
func NewIOSClient() IOSClient {
	return IOSClient{}
}
