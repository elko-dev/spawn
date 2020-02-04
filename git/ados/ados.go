package ados

import (
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// Repository struct to leverage Azure DevOps
type Repository struct {
	prompt Prompt
}

// Prompt for user info
type Prompt interface {
	forOrganization() (string, error)
	forGitToken() (string, error)
}

// CreateGitRepository action to create an ADOS repo
func (ados Repository) CreateGitRepository(repositoryName string, templateURL string, platformToken string) error {

	organization, err := ados.prompt.forOrganization()

	if err != nil {
		log.Debug("Error creating organization")
		return err
	}

	organizationURL := "https://dev.azure.com/" + organization

	gitToken, err := ados.prompt.forGitToken()

	if err != nil {
		return err
	}
	contextLogger := log.WithFields(log.Fields{
		"repositoryName":  repositoryName,
		"templateURL":     templateURL,
		"organizationURL": organizationURL,
	})

	connection := azuredevops.NewPatConnection(organizationURL, gitToken)
	ctx := context.Background()
	coreClient, err := core.NewClient(ctx, connection)
	
	if err != nil {
		return err
	}
	contextLogger.Debug("created ados context")
	description := "test"
	var testID = uuid.New()

	adosProject := core.TeamProject{
		Name:        &repositoryName,
		Description: &description,
		Visibility:  &core.ProjectVisibilityValues.Private,
		Capabilities: &map[string]map[string]string{
			"versioncontrol":  {"sourceControlType": "Git"},
			"processTemplate": {"templateTypeId": testID.String()},
		}}
	projectArgs := core.QueueCreateProjectArgs{ProjectToCreate: &adosProject}

	_, err = coreClient.QueueCreateProject(ctx, projectArgs)
	if err != nil {
		return err
	}
	contextLogger.Debug("ados project queued")

	return nil
}

// NewRepository init method
func NewRepository(prompt Prompt) Repository {
	return Repository{prompt}
}
