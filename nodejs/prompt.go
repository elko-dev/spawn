package nodejs

import (
	"github.com/elko-dev/spawn/constants"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// Prompts stuct to prompt user for projectName
type Prompts struct {
}

func (prompts Prompts) forAppName() (string, error) {
	selection, err := selectProjectName()
	return selection, err
}

func (prompts Prompts) forPlatform() (string, error) {
	_, selection, err := selectPlatform()
	return selection, err
}

func (prompts Prompts) forVersionControl() (string, error) {
	_, versionControl, err := selectVersionControl()
	return versionControl, err
}

func selectVersionControl() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Version Control",
		Items: []string{constants.Gitlab, constants.ADOS},
	}

	return prompt.Run()
}

//TODO: Remove this
func selectPlatform() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Platform",
		Items: []string{constants.HerokuPlatform, constants.AzureFunctions},
	}

	return prompt.Run()
}

func selectProjectName() (string, error) {
	projectValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
				is.Alphanumeric,
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "NodeJS Project Name",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
