package platform

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forPlatformType() (string, error) {
	_, selection, err := selectPlatform()
	return selection, err
}

func selectPlatform() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Platform",
		Items: []string{constants.ExpressHerokuPlatform,
			constants.GraphQLHerokuPlatform,
			constants.AzureFunctions},
	}

	return prompt.Run()
}

func NewPrompts() Prompt {
	return Prompts{}
}
