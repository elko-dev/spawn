package applicationtype

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/manifoldco/promptui"
)

// Prompts struct to use to prompt user for selections
type Prompts struct {
}

// ForType prompts user to select Application Type
func (prompts Prompts) ForType() (string, error) {
	_, selection, err := selectApplicationType()
	return selection, err
}

func selectApplicationType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Application Type",
		Items: []string{constants.WebApplicationType, constants.MobileApplicationType, constants.AzureFunctions},
	}

	return prompt.Run()
}

// NewPrompts init
func NewPrompts() Prompts {
	return Prompts{}
}
