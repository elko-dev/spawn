package appcenter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Client struct to send http request
type Client struct {
	baseURL       string
	client        *http.Client
	authorization string
}

// SendRequest sends http request
func (client *Client) SendRequest(request *http.Request) (response *http.Response, err error) {
	// TODO: add retry logic
	request.Header.Add("X-API-Token", client.authorization)
	request.Header.Add("accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	resp, err := client.client.Do(request)
	if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
		unwrappedError := ErrorResponse{}
		unwrapError(resp, unwrappedError)

		log.WithFields(log.Fields{
			"responseCode": resp.StatusCode,
			"errorCode":    unwrappedError.Error.code,
			"message":      unwrappedError.Error.message,
		}).Error("Sending request to App Center")

		return resp, errors.New("Error received when calling AppCenter API")
	}
	return resp, err
}

// Send http request to appcenter api
func (client *Client) Send(ctx context.Context,
	httpMethod string,
	body io.Reader) (response *http.Response, err error) {

	req, err := http.NewRequest(httpMethod, client.baseURL, body)

	log.WithFields(log.Fields{
		"httpMethod": httpMethod,
		"url":        client.baseURL,
	}).Debug("Sending request to App Center")

	resp, err := client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// UnmarshalBody unmarshalls response body
func (client *Client) UnmarshalBody(response *http.Response, v interface{}) (err error) {
	if response != nil && response.Body != nil {
		var err error
		defer func() {
			if closeError := response.Body.Close(); closeError != nil {
				err = closeError
			}
		}()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
		return json.Unmarshal(body, &v)
	}
	return nil
}

func unwrapError(response *http.Response, target interface{}) error {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	println(bodyString)
	if err != nil {
		println(err)
		return err
	}
	return json.Unmarshal(bodyBytes, target)
}

// NewClient init func
func NewClient(connection *Connection, baseURL string) *Client {
	client := &http.Client{}
	return &Client{
		baseURL:       baseURL,
		client:        client,
		authorization: connection.AuthorizationToken,
	}
}
