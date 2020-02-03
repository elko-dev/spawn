package functions

// FunctionsType struct to create an Azure Functions type
type FunctionsType struct {
}

// Create sets up a new application
func (function FunctionsType) Create() error {
	return nil
}

// NewFunctionsType init function
func NewFunctionsType() FunctionsType {
	return FunctionsType{}
}
