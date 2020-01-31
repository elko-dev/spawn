package prompt

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

// HerokuCommand struct to retrieve user commands
type HerokuCommand struct {
}

// Platform test
func (platform HerokuCommand) Platform() (string, string, error) {
	token, err := selectPlatformToken()
	if err != nil {
		return "", "",  err
	}

	teamName, err := selectHerokuTeamName()
	if err != nil {
		return "", "",err
	}
	return token, teamName, nil
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
