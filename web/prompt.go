package web

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/manifoldco/promptui"
)

// Prompts struct to
type Prompts struct {
}

// ForServerType prompts user for server type
func (prompts Prompts) ForServerType() (string, error) {
	_, selection, err := selectServerType()
	return selection, err
}

func selectServerType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Backend Language",
		Items: []string{constants.NodeServerType},
	}

	return prompt.Run()
}

// ForClientType prompts user for client type
func (prompts Prompts) ForClientType(applicationType string) (string, error) {
	clientLanguageType := getClientLangaugeSelections(applicationType)
	_, selection, err := selectClientLanguageTypes(clientLanguageType)
	return selection, err
}

func selectClientLanguageTypes(selections []string) (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Client Language",
		Items: selections,
	}

	return prompt.Run()
}

//Sorry for the term but...helper function
func getClientLangaugeSelections(applicationType string) []string {
	if applicationType == constants.WebApplicationType {
		return []string{constants.ReactClientLanguageType}
	}
	if applicationType == constants.MobileApplicationType {
		return []string{"React Native"}
	}
	//TODO: This seems to be a bit of a hack; need to rethink the interface
	if applicationType == constants.AzureFunctions {
		return []string{"None"}
	}
	return make([]string, 0, 0)
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
