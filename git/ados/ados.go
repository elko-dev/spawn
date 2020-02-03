package ados

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// ADOSRepository struct to leverage Azure DevOps
type Repository struct {
	prompt Prompt
}

// ADOSPrompt for user info
type Prompt interface {
	forOrganization() (string, error)
	forGitToken() (string, error)
}

// CreateGitRepository action to create an ADOS repo
func (ados Repository) CreateGitRepository(repositoryName string, url string, platformToken string) error {
	organization, err := ados.prompt.forOrganization()

	if err != nil {
		return err
	}

	organizationURL := "https://dev.azure.com/" + organization

	gitToken, err := ados.prompt.forGitToken()

	if err != nil {
		return err
	}

	connection := azuredevops.NewPatConnection(organizationURL, gitToken)
	ctx := context.Background()
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		return err
	}

	repoName := "repositoryName"
	adosProject := core.TeamProject{Name: &repoName}
	projectArgs := core.QueueCreateProjectArgs{ProjectToCreate: &adosProject}
	_, err = coreClient.QueueCreateProject(ctx, projectArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewRepository init method
func NewRepository(prompt Prompt) Repository {
	return Repository{prompt}
}
