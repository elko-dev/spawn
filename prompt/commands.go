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
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "Gitlab Access Token",
		Validate: accessTokenValidate,
		Mask:     '*',
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
		Mask:     '*',
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

// GitlabGroupID prompts user for Gitlab group id to add repo
func GitlabGroupID() (string, error) {
	//TODO: Add custom id validation if we stick with group id instead of name
	gitlabGroupValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	projectPrompt := promptui.Prompt{
		Label:    "Gitlab Group Id",
		Validate: gitlabGroupValidate,
	}

	return projectPrompt.Run()
}
