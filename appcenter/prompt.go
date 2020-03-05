package appcenter

import (
	"strings"

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
func (prompt Prompts) forAuthSecretPath() (string, error) {
	secretPath, err := firebaseSecretPath()

	if err != nil {
		return "", err
	}
	return secretPath, nil
}

func firebaseSecretPath() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "Path to Firebase Secret file; relative to spawn (include file name).",
		Validate: accessTokenValidate,
	}

	return accessTokenPrompt.Run()
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

func (prompt Prompts) forMembers() ([]string, error) {
	memberString, err := members()

	if err != nil {
		return make([]string, 0), err
	}
	return strings.Split(memberString, ","), nil
}
func members() (string, error) {
	accessTokenValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	accessTokenPrompt := promptui.Prompt{
		Label:    "Distribution member emails (comma separated list)",
		Validate: accessTokenValidate,
	}

	return accessTokenPrompt.Run()
}

// NewPrompts init
func NewPrompts() Prompt {
	return Prompts{}
}
