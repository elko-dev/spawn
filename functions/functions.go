package functions

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/prompt"
	"github.com/elko-dev/spawn/web"
)

// FunctionsType struct to create an Azure Functions type
type FunctionsType struct {
	function applications.App
}

// Create sets up a new application
func (function FunctionsType) Create(action web.SpawnAction, userCommands prompt.UserSelections) error {

	return nil
}

// NewFunctionsType init function
func NewFunctionsType(function applications.App) FunctionsType {
	return FunctionsType{function}
}
