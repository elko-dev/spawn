package appcenter

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forOrganization() (string, error) {
	organization, err := appcenterOrganization()

	if err != nil {
		return "", err
	}
	return organization, nil
}

func appcenterOrganization() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "AppCenter Organization",
		Validate: accessTokenValidate,
	}

	return accessTokenPrompt.Run()
}
func (prompt Prompts) forToken() (string, error) {
	token, err := appcenterToken()

	if err != nil {
		return "", err
	}
	return token, nil
}

func appcenterToken() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "AppCenter Access Token",
		Validate: accessTokenValidate,
		Mask:     '*',
	}

	return accessTokenPrompt.Run()
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}