package builds

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elko-dev/spawn/appcenter/api"
)

const (
	repoPathURL        = "apps/{owner_name}/{app_name}/repo_config"
	buildPathURL       = "apps/{owner_name}/{app_name}/branches/{branch}/builds"
	configBuildURL     = "apps/{owner_name}/{app_name}/branches/{branch}/config"
	orgReplaceParam    = "{owner_name}"
	appReplaceParam    = "{app_name}"
	branchReplaceParam = "{branch}"
)

// Client interface for appcenter
type Client interface {
	ConfigureRepo(context context.Context, args *RepoConfigArgs, orgName string, projectName string) (*RepoConfigResponse, error)
	Build(context context.Context, args *BuildArgs, orgName string, projectName string) (*BuildResponse, error)
	ConfigureBuild(context context.Context, args *ConfigArgs, orgName string, projectName string) (*BuildConfigResponse, error)
}

// ClientImpl implementation of Client interface
type ClientImpl struct {
	Client api.Client
}

// ConfigureBuild to configure Build
func (client ClientImpl) ConfigureBuild(context context.Context, args *ConfigArgs, orgName string, projectName string) (*BuildConfigResponse, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), createConfigBuildURL(orgName, projectName, "master"))
	if err != nil {
		return nil, err
	}

	var responseValue BuildConfigResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Build project
func (client ClientImpl) Build(context context.Context, args *BuildArgs, orgName string, projectName string) (*BuildResponse, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), createBuildURL(orgName, projectName, "master"))
	if err != nil {
		return nil, err
	}

	var responseValue BuildResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// ConfigureRepo configures appcenter repository
func (client ClientImpl) ConfigureRepo(context context.Context, args *RepoConfigArgs, orgName string, projectName string) (*RepoConfigResponse, error) {
	body, marshalErr := json.Marshal(args)
	if marshalErr != nil {
		return nil, marshalErr
	}
	resp, err := client.Client.Send(context, http.MethodPost, bytes.NewReader(body), createRepoURL(orgName, projectName))
	if err != nil {
		return nil, err
	}

	var responseValue RepoConfigResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

func createRepoURL(orgName string, projectName string) string {
	return strings.Replace(strings.Replace(repoPathURL, orgReplaceParam, orgName, 1), appReplaceParam, projectName, 1)
}

func createBuildURL(orgName string, projectName string, branch string) string {
	orgReplacedString := strings.Replace(buildPathURL, orgReplaceParam, orgName, 1)
	appReplactedString := strings.Replace(orgReplacedString, appReplaceParam, projectName, 1)
	branchReplacedString := strings.Replace(appReplactedString, branchReplaceParam, branch, 1)
	return branchReplacedString
}
func createConfigBuildURL(orgName string, projectName string, branch string) string {
	orgReplacedString := strings.Replace(configBuildURL, orgReplaceParam, orgName, 1)
	appReplactedString := strings.Replace(orgReplacedString, appReplaceParam, projectName, 1)
	branchReplacedString := strings.Replace(appReplactedString, branchReplaceParam, branch, 1)
	return branchReplacedString
}

// NewClient init
func NewClient(connection *api.Connection) Client {
	client := connection.GetClient()
	return &ClientImpl{
		Client: *client,
	}
}
