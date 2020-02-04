package azurefunctions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/platform"
)

type Factory struct {
}

func (factory Factory) Create(applicationType string) (applications.PlatformRepository, error) {
	println("Functions platform")
	return NewAzureFunctions(), nil
}
func NewFactory() platform.FunctionsPlatformFactory {
	return Factory{}
}
