package ados

import (
	"context"
	"time"

	// "github.com/google/uuid"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/git/local"
	log "github.com/sirupsen/logrus"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// Repository struct to leverage Azure DevOps
type Repository struct {
	prompt Prompt
	Git    local.Git
}

// Prompt for user info
type Prompt interface {
	forOrganization() (string, error)
}

// CreateGitRepository action to create an ADOS repo
func (ados Repository) CreateGitRepository(repositoryName string, templateURL string, platformToken string, replacements map[string]string) (applications.GitResult, error) {

	organization, err := ados.prompt.forOrganization()

	if err != nil {
		log.Debug("Error creating organization")
		return applications.GitResult{}, err
	}

	organizationURL := "https://dev.azure.com/" + organization

	if err != nil {
		return applications.GitResult{}, err
	}
	contextLogger := log.WithFields(log.Fields{
		"repositoryName":  repositoryName,
		"templateURL":     templateURL,
		"organizationURL": organizationURL,
	})

	connection := azuredevops.NewPatConnection(organizationURL, platformToken)
	ctx := context.Background()
	coreClient, err := core.NewClient(ctx, connection)

	if err != nil {
		return applications.GitResult{}, err
	}

	contextLogger.Debug("created ados context")
	description := "test"
	// IT IS WORTH NOTING....
	// This is some random default ID that microsoft doesn't seem to want to bother documenting what it is
	// nor how to generate one.  I pulled this from the below link:
	// https://developercommunity.visualstudio.com/content/problem/176992/unable-to-create-project-from-rest.html
	var templateID = "27450541-8e31-4150-9947-dc59f998fc01"

	adosProject := core.TeamProject{
		Name:        &repositoryName,
		Description: &description,
		Visibility:  &core.ProjectVisibilityValues.Private,
		Capabilities: &map[string]map[string]string{
			"versioncontrol":  {"sourceControlType": "Git"},
			"processTemplate": {"templateTypeId": templateID},
		}}
	projectArgs := core.QueueCreateProjectArgs{ProjectToCreate: &adosProject}

	response, err := coreClient.QueueCreateProject(ctx, projectArgs)
	if err != nil {
		contextLogger.Debug("failed to create ados project")
		return applications.GitResult{}, err
	}
	contextLogger.Debug(
		"ados project queued with response ",
		*response.Url,
	)

	adosGitRepoURL := "https://dev.azure.com/" + organization + "/" + repositoryName + "/_git/" + repositoryName

	log.WithFields(log.Fields{}).Info("Waiting for azure repo to create....")
	time.Sleep(5 * time.Second)

	return ados.Git.DuplicateRepo(templateURL, platformToken, repositoryName, adosGitRepoURL, replacements)
}

// NewRepository init method
func NewRepository(prompt Prompt) Repository {
	return Repository{prompt, local.NewLocal()}
}
