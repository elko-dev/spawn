package git

import (
	"context"

	"github.com/elko-dev/spawn/git/api"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// ADOSRepository struct to leverage Azure DevOps
type ADOSRepository struct {
}

// CreateGitRepository action to create an ADOS repo
func (ados ADOSRepository) CreateGitRepository(repositoryName string, gitToken string, platformToken string, url string) (api.GitRepository, error) {
	organizationURL := "https://dev.azure.com/elko-playground"
	connection := azuredevops.NewPatConnection(organizationURL, gitToken)
	ctx := context.Background()
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		return api.GitRepository{}, err
	}

	adosProject := core.TeamProject{Name: &repositoryName}
	projectArgs := core.QueueCreateProjectArgs{ProjectToCreate: &adosProject}
	response, err := coreClient.QueueCreateProject(ctx, projectArgs)
	if err != nil {
		return api.GitRepository{}, err
	}
	return api.GitRepository{Name: repositoryName, URL: *response.Url}, nil
}

// NewADOSRepository init method
func NewADOSRepository(git Git) ADOSRepository {

	return ADOSRepository{}
}
