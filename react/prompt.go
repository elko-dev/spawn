package react

import (
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

func selectProjectName() (string, error) {
	projectValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
				is.Alphanumeric,
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "React Project Name",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
