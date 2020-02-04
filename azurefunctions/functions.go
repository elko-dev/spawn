package azurefunctions

import "github.com/elko-dev/spawn/applications"

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

func NewAzureFunctions() applications.PlatformRepository {
	return AzureFunctions{}
}
