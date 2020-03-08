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
func (client IOSClient) Create(firebaseProjectID string, request IOSRequest) (applications.IOSApp, error) {

	req, err := createIOSRequest(firebaseProjectID, request)
	if err != nil {
		return applications.IOSApp{}, err
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return applications.IOSApp{}, err
	}

	if !spawnhttp.IsSuccessStatusCode(resp.StatusCode) {
		log.WithFields(log.Fields{
			"response":          resp,
			"firebaseProjectID": firebaseProjectID,
		}).Error("Error creating IOS App")

		return applications.IOSApp{}, errors.New("Received error creating ios project with status ")
	}

	iosResponse := IOSResponse{}
	err = spawnhttp.MarshalResponse(resp, &iosResponse)

	if err != nil {
		return applications.IOSApp{}, err
	}

	response := applications.IOSApp{
		ID:   iosResponse.ID,
		Name: iosResponse.Name,
	}
	log.WithFields(log.Fields{
		"rawResponse": resp,
		"iosResponse": response,
	}).Debug("Successfully created IOS Project")

	return response, nil
}

func createIOSRequest(firebaseProjectID string, request IOSRequest) (*http.Request, error) {
	requestURL := strings.Replace(iosURL, ":id", firebaseProjectID, 1)
	return spawnhttp.CreateRequest(requestURL, request)
}

// NewIOSClient init
func NewIOSClient() IOSClient {
	return IOSClient{}
}
