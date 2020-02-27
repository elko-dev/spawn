package organization

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/elko-dev/spawn/appcenter/api"
)

const (
	orgPathURL = "orgs"
)

// Client interface for appcenter
type Client interface {
	CreateOrganization(context context.Context, args *CreateOrganizationArgs) (*Organization, error)
}

// ClientImpl implementation of Client interface
type ClientImpl struct {
	Client api.Client
}

// CreateOrganization creates organization
func (client ClientImpl) CreateOrganization(context context.Context, args *CreateOrganizationArgs) (*Organization, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), orgPathURL)
	if err != nil {
		return nil, err
	}

	var responseValue Organization
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// NewClient init
func NewClient(connection *api.Connection) Client {
	client := connection.GetClient()
	return &ClientImpl{
		Client: *client,
	}
}
