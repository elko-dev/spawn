package reactnative

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// Prompts stuct to prompt user for projectName
type Prompts struct {
}

func (prompts Prompts) forIncludingPlatform() (bool, error) {
	prompt := promptui.Prompt{
		Label:     "Include Firebase Platform?",
		IsConfirm: true,
	}

	confirmation, _ := prompt.Run()

	return isTrueOrFalse(confirmation), nil
}

func isTrueOrFalse(confirmation string) bool {
	if strings.ToLower(confirmation) == "y" {
		return true
	}

	return false
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
		Label:    "React Native Project Name",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
