package prompt

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// GitlabAccessToken prompts user for gitlab token
func GitlabAccessToken() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
				is.Alphanumeric,
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "Gitlab Access Token",
		Validate: accessTokenValidate,
	}

	return accessTokenPrompt.Run()
}

// DeployAccessToken prompts user for heroku token
func DeployAccessToken() (string, error) {
	deployTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	deployTokenPrompt := promptui.Prompt{
		Label:    "Heroku Access Token",
		Validate: deployTokenValidate,
	}

	return deployTokenPrompt.Run()
}

// ProjectName prompts user for project name
func ProjectName() (string, error) {
	projectValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
				is.Alphanumeric,
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "Project Name",
		Validate: projectValidate,
	}

	return projectPrompt.Run()
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
