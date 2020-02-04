package azurefunctions

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/pipelines"
)

// AzureFunctions stuct to create function
type AzureFunctions struct {
	organization string
	token        string
	projectName  string
}

// Create AzureFunction
func (f AzureFunctions) Create() error {

	organizationURL := "https://dev.azure.com/" + f.organization

	contextLogger := log.WithFields(log.Fields{
		"organization": f.organization,
		"projectName":  f.projectName,
	})

	connection := azuredevops.NewPatConnection(organizationURL, f.token)
	ctx := context.Background()
	coreClient := pipelines.NewClient(ctx, connection)

	contextLogger.Debug("Created pipelines client")
	params := pipelines.CreatePipelineParameters{
		Configuration: &pipelines.CreatePipelineConfigurationParameters{
			Type: &pipelines.ConfigurationTypeValues.Yaml,
		},
	}
	args := pipelines.CreatePipelineArgs{
		Project:         &f.projectName,
		InputParameters: &params,
	}
	pipelineResponse, err := coreClient.CreatePipeline(ctx, args)
	if err != nil {
		contextLogger.Debug("Error creating pipeline")
		return err
	}

	_, err = coreClient.RunPipeline(ctx, pipelines.RunPipelineArgs{
		Project:    &f.projectName,
		PipelineId: pipelineResponse.Id,
	})
	return err
}

// GetToken for ADOS
func (f AzureFunctions) GetToken() string {
	return f.token
}

// GetPlatformType for azure functions
func (f AzureFunctions) GetPlatformType() string {
	return constants.AzureFunctions
}

// NewAzureFunctions init
func NewAzureFunctions(organization string, token string, projectName string) applications.PlatformRepository {
	return AzureFunctions{organization, token, projectName}
}
