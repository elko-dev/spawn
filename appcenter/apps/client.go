package apps

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/appcenter/api"
)

const (
	orgPathURL   = "orgs/{org_name}/apps"
	replaceParam = "{org_name}"
)

// Client interface for appcenter
type Client interface {
	CreateApp(context context.Context, args *CreateAppArgs, orgName string) (*App, error)
}

// ClientImpl implementation of Client interface
type ClientImpl struct {
	Client api.Client
}

// CreateApp creates an application
func (client ClientImpl) CreateApp(context context.Context, args *CreateAppArgs, orgName string) (*App, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var responseValue App
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

func createOrganizationURL(orgName string) string {
	return strings.Replace(orgPathURL, replaceParam, orgName, 1)
}

// NewClient init
func NewClient(connection *api.Connection, orgName string) Client {
	fullURL := createOrganizationURL(orgName)
	client := connection.GetClientByAPIURL(fullURL)
	return &ClientImpl{
		Client: *client,
	}
}
