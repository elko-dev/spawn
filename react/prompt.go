package react

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// Prompts stuct to prompt user for projectName
type Prompts struct {
}

func (prompts Prompts) forWebId() (string, error) {
	selection, err := selectWebId()
	return selection, err
}

func selectWebId() (string, error) {
	projectValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "Firebase WebId (this can be found in firebase)",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
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
				is.LowerCase,
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
