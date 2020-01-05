package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	gitlabProjectURL = "https://gitlab.com/api/v4/projects"
)

// GitlabHTTP struct to leverage Gitlab
type GitlabHTTP struct {
}

// GitRepository struct containing information about git repository
type GitRepository struct {
	Name string `json:"name"`
	URL  string `json:"http_url_to_repo"`
	ID   string `json:"id"`
}

// PostGitRepository Creates Git Repository
func (rest GitlabHTTP) PostGitRepository(repositoryName string, accessToken string) (GitRepository, error) {
	var projectRequest = []byte(`{"name":"` + repositoryName + `"}`)
	req, err := http.NewRequest("POST", gitlabProjectURL, bytes.NewBuffer(projectRequest))
	req.Header.Set("PRIVATE-TOKEN", accessToken)
	req.Header.Set("Content-Type", "application/json")

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
	//todo: parse my body
	return response, errors.New("Error creating gitlab repo")
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
