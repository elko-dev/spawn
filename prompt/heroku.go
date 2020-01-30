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
	token, err := PlatformToken()
	if err != nil {
		return "","", err
	}

	teamName, err := HerokuTeamName()
	if err != nil {
		return  "","", err
	}
	return token, teamName, nil
}

// PlatformToken prompts user for heroku token
func PlatformToken() (string, error) {
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

// HerokuTeamName prompts user for heroku team name
func HerokuTeamName() (string, error) {
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
