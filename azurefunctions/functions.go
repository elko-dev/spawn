package azurefunctions

import "github.com/elko-dev/spawn/applications"

import "github.com/elko-dev/spawn/constants"

type AzureFunctions struct {
}

func (f AzureFunctions) Create() error {
	println("Azure function executing")
	return nil
}

func (f AzureFunctions) GetToken() string {
	println("Azure function executing")
	return ""
}

func (f AzureFunctions) GetPlatformType() string {
	return constants.AzureFunctions
}

func NewAzureFunctions() applications.PlatformRepository {
	return AzureFunctions{}
}
