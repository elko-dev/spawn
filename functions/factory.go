package functions

import "github.com/elko-dev/spawn/web"

// Factory to create Functions App
type Factory struct {
	nodeJsFactory web.AppFactory
}

// Create returns a FunctionType
func (factory Factory) Create() FunctionsType {
	nodeJs := factory.nodeJsFactory.Create()

	return NewFunctionsType(nodeJs)
}

//NewFactory init function
func NewFactory(nodeJsFactory web.AppFactory) Factory {
	return Factory{nodeJsFactory}
}
