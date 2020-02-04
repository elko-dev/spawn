package azurefunctions

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forToken() (string, error) {
	token, err := adosToken()

	if err != nil {
		return "", err
	}
	return token, nil
}

func adosToken() (string, error) {
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
func (prompt Prompts) forOrganization() (string, error) {
	organization, err := adosOrganization()

	if err != nil {
		return "", err
	}
	return organization, nil
}

func adosOrganization() (string, error) {
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

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
