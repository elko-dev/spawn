package accounts

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/appcenter/api"
)

const (
	distributionURL        = "orgs/{org_name}/distribution_groups"
	appsForDistributionURL = "orgs/{org_name}/distribution_groups/{distribution_group_name}/apps"
	memberInviteURL        = "orgs/{org_name}/distribution_groups/{distribution_group_name}/members"
	orgReplaceParam        = "{org_name}"
	distributionGroupParam = "{distribution_group_name}"
)

// Client to manage account details
type Client interface {
	CreateDistributionGroup(context context.Context, args *DistributionGroupArg, orgName *string) (*DistributionGroupResponse, error)
	CreateAppsDistributionGroup(context context.Context, args *AppsForDistributionArg, orgName *string, distributionName *string) error
	AddMemberToDistribution(context context.Context, args *AddMemberArgs, orgName *string, distributionName *string) error
}

// ClientImpl implementation of Client interface
type ClientImpl struct {
	Client api.Client
}

// CreateAppsDistributionGroup to add app to distribution group
func (client ClientImpl) CreateAppsDistributionGroup(context context.Context, args *AppsForDistributionArg, orgName *string, distributionName *string) error {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return marshalErr
	}
	_, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), createAppsForDistributionURL(*orgName, *distributionName))

	return err
}

// AddMemberToDistribution to add memebers to distribution group
func (client ClientImpl) AddMemberToDistribution(context context.Context, args *AddMemberArgs, orgName *string, distributionName *string) error {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return marshalErr
	}
	_, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), addMemberForDistributionURL(*orgName, *distributionName))

	return err
}

// CreateDistributionGroup for organization
func (client ClientImpl) CreateDistributionGroup(context context.Context, args *DistributionGroupArg, orgName *string) (*DistributionGroupResponse, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), createDistributionURL(*orgName))
	if err != nil {
		return nil, err
	}

	var responseValue DistributionGroupResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

func createDistributionURL(orgName string) string {
	return strings.Replace(distributionURL, orgReplaceParam, orgName, 1)
}
func createAppsForDistributionURL(distributionName string, orgName string) string {
	orgReplaced := strings.Replace(appsForDistributionURL, orgReplaceParam, orgName, 1)
	return strings.Replace(orgReplaced, distributionGroupParam, distributionName, 1)
}
func addMemberForDistributionURL(distributionName string, orgName string) string {
	orgReplaced := strings.Replace(memberInviteURL, orgReplaceParam, orgName, 1)
	return strings.Replace(orgReplaced, distributionGroupParam, distributionName, 1)
}

// NewClient init
func NewClient(connection *api.Connection) Client {
	client := connection.GetClient()
	return &ClientImpl{
		Client: *client,
	}
}
