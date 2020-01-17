package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elko-dev/spawn/prompt"
	"io/ioutil"
	"net/http"
)

const (
	gitlabProjectURL     = "https://gitlab.com/api/v4/projects"
	gitlabEnvironmentURL = "https://gitlab.com/api/v4/projects/%v/variables"
)

// GitlabHTTP struct to leverage Gitlab
type GitlabHTTP struct {
	// client http.Client
}

// GitRepository struct containing information about git repository
type GitRepository struct {
	Name string      `json:"name"`
	URL  string      `json:"http_url_to_repo"`
	ID   json.Number `json:"id,Number"`
}

// AddEnvironmentVariables to project
func (rest GitlabHTTP) AddEnvironmentVariables(deployToken string, projectID string, accessToken string) error {
	environmentRequest := []byte(`{
		"key": "HEROKU_API_KEY",
		"value": "` + deployToken + `",
		"protected": true
	}`)
	url := fmt.Sprintf(gitlabEnvironmentURL, projectID)
	req, err := createPostRequest(accessToken, url, environmentRequest)

	if err != nil {
		println("Error adding environment variables")
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if isSuccessStatusCode(resp.StatusCode) {
		return nil
	}

	if resp.StatusCode == 401 {
		fmt.Println("Received unauthorized from Gitlab")
		return errors.New("Unauthorized")
	}

	println("Failed to add environment variables")
	println(resp.StatusCode)

	return errors.New("Bad Request")
}

// PostGitRepository Creates Git Repository
func (rest GitlabHTTP) PostGitRepository(repositoryName string, accessToken string) (GitRepository, error) {
	group, err := prompt.GitlabGroupID()
	if err != nil {
		println("Error retrieving Gitlab Group name")
		return GitRepository{}, err
	}
	var projectRequest = createProjectRequest(repositoryName, group)
	req, err := createPostRequest(accessToken, gitlabProjectURL, projectRequest)

	client := &http.Client{}
	resp, err := client.Do(req)

	response := GitRepository{}

	if err != nil {
		println("Error creating Gitlab Repository; aborting...")
		return response, err
	}
	defer resp.Body.Close()

	if isSuccessStatusCode(resp.StatusCode) {
		parseGitlabResponse(resp, &response)
		return response, nil
	}

	if resp.StatusCode == 401 {
		fmt.Println("Received unauthorized from Gitlab")
		return response, errors.New("Unauthorized")
	}
	println("Failed to create gitlab repository")
	println(resp.StatusCode)
	return response, errors.New("Error creating gitlab repo")
}

func createProjectRequest(respositoryName string, group string) []byte {
	return []byte(`{"path":"` + respositoryName + `", "namespace_id": ` + group + `}`)
}

func createPostRequest(accessToken string, url string, request []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	req.Header.Set("PRIVATE-TOKEN", accessToken)
	req.Header.Set("Content-Type", "application/json")
	return req, err
}

func isSuccessStatusCode(statusCode int) bool {
	return statusCode == 201 || statusCode == 200
}

func isUnauthorized(statusCode int) bool {
	return statusCode == 401
}

func parseGitlabResponse(response *http.Response, target interface{}) error {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err)
		return err
	}
	return json.Unmarshal(bodyBytes, target)
}

// // NewGitlabHTTP init function
// func NewGitlabHTTP() git.HTTP {
// 	client := &http.Client{}

// 	return GitlabHTTP{client: client}
// }
