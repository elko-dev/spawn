package nodejs

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// Prompts stuct to prompt user for projectName
type Prompts struct {
}

func (prompts Prompts) ForAppName() (string, error) {
	selection, err := selectProjectName()
	return selection, err
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
