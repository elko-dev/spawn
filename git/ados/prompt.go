package ados

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forOrganization() (string, error) {
	organization, err := gitOrganization()

	if err != nil {
		return "", err
	}
	return organization, nil
}
func (prompt Prompts) forGitToken() (string, error) {
	token, err := gitAccessToken()

	if err != nil {
		return "", err
	}
	return token, nil
}

func gitAccessToken() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "ADOS Access Token",
		Validate: accessTokenValidate,
		Mask:     '*',
	}

	return accessTokenPrompt.Run()
}

func gitOrganization() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "ADOS Organization",
		Validate: accessTokenValidate,
	}

	return accessTokenPrompt.Run()
}

// NewPrompts init func
func NewPrompts() Prompts {
	return Prompts{}
}
