package gitlab

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	prompt Prompt
}

//TODO: Factor out ID to be used by multiple git repos
// GitRepository struct containing information about git repository
type GitRepository struct {
	Name string      `json:"name"`
	URL  string      `json:"http_url_to_repo"`
	ID   json.Number `json:"id,Number"`
}

// AddEnvironmentVariables to project
func (rest GitlabHTTP) AddEnvironmentVariables(deployToken string, projectID string, gitToken string) error {
	environmentRequest := []byte(`{
		"key": "HEROKU_API_KEY",
		"value": "` + deployToken + `",
		"protected": true
	}`)
	url := fmt.Sprintf(gitlabEnvironmentURL, projectID)
	req, err := createPostRequest(gitToken, url, environmentRequest)

	if err != nil {
		println("Error adding environment variables")
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if isSuccessStatusCode(resp.StatusCode) {
		return nil
	}

	if isUnauthorized(resp.StatusCode) {
		fmt.Println("Received unauthorized from Gitlab")
		return errors.New("Unauthorized")
	}

	println("Failed to add environment variables")
	println(resp.StatusCode)

	return errors.New("Bad Request")
}

// PostGitRepository Creates Git Repository
func (rest GitlabHTTP) PostGitRepository(repositoryName string, gitToken string) (GitRepository, error) {
	group, err := rest.prompt.forGroupId()
	if err != nil {
		println("Error retrieving Gitlab Group name")
		return GitRepository{}, err
	}
	var projectRequest = createProjectRequest(repositoryName, group)
	req, err := createPostRequest(gitToken, gitlabProjectURL, projectRequest)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		println("Error creating Gitlab Repository; aborting...")
		return GitRepository{}, err
	}
	defer resp.Body.Close()

	if isSuccessStatusCode(resp.StatusCode) {
		response := GitRepository{}
		parseGitlabResponse(resp, &response)
		return response, nil
	}

	if isUnauthorized(resp.StatusCode) {
		fmt.Println("Received unauthorized from Gitlab")
		return GitRepository{}, errors.New("Unauthorized")
	}

	println("Failed to create gitlab repository")
	println(resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	return GitRepository{}, errors.New("Error creating gitlab repo")
}

func createProjectRequest(respositoryName string, group string) []byte {
	return []byte(`{"path":"` + respositoryName + `",` + `"visibility":"private",` + ` "namespace_id": ` + group + `}`)
}

func createPostRequest(gitToken string, url string, request []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(request))
	req.Header.Set("PRIVATE-TOKEN", gitToken)
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

// NewGitlabHTTP init
func NewGitlabHTTP(prompt Prompt) GitlabHTTP {
	return GitlabHTTP{prompt}
}
