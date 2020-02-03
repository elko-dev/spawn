package herokuplatform

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
	"strings"
)

// Prompts user for platform fields
type Prompts struct {
}

func (prompts Prompts) forEnvironments() ([]string, error) {
	envString, err := selectEnvironments()
	if err != nil {
		return nil, err
	}

	return strings.Split(envString, ","), nil
}
func (prompts Prompts) forHerokuTeamName() (string, error) {
	return selectHerokuTeamName()
}
func (prompts Prompts) forPlatformToken() (string, error) {
	return selectPlatformToken()
}

func selectPlatformToken() (string, error) {
	deployTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	deployTokenPrompt := promptui.Prompt{
		Label:    "Heroku Access Token",
		Validate: deployTokenValidate,
		Mask:     '*',
	}

	return deployTokenPrompt.Run()
}

func selectHerokuTeamName() (string, error) {
	projectValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "Heroku Team Name",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
}
func selectEnvironments() (string, error) {
	envValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "Environments (comma separated list)",
		Validate: envValidate,
	}

	return projectPrompt.Run()
}

func NewPrompts() Prompt {
	return Prompts{}
}