package functions

import (
	"github.com/elko-dev/spawn/applications"
)

// FunctionsType struct to create an Azure Functions type
type FunctionsType struct {
	project applications.Project
}

// Create sets up a new application
func (function FunctionsType) Create() error {
	return function.project.Create()
}

// NewFunctionsType init function
func NewFunctionsType(project applications.Project) FunctionsType {
	return FunctionsType{project}
}
