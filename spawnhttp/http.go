package spawnhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func IsSuccessStatusCode(statusCode int) bool {
	return statusCode == 201 || statusCode == 200
}

func MarshalResponse(response *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &target)
	if err != nil {
		return err
	}

	return nil
}

// CreateRequest http
func CreateRequest(url string, request interface{}) (*http.Request, error) {
	bytesRepresentation, err := json.Marshal(request)
	if err != nil {
		return &http.Request{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
