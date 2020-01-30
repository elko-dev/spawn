package prompt

import (
	"github.com/elko-dev/spawn/constants"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

// GitPrompts specifies methods to retrieve git configuration
type GitPrompts struct {
}

// Token allows user to input token
func (gitlab GitPrompts) Token() (string, error) {
	token, err := gitAccessToken()

	if err != nil {
		return "", err
	}
	return token, nil
}

// Repository allows user to select git repo
func (gitlab GitPrompts) Repository() (string, error) {
	_, gitRepo, err := selectGitRepository()

	if err != nil {
		return "", err
	}
	return gitRepo, nil
}

func selectGitRepository() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Git Repository",
		Items: []string{constants.ADOS, constants.Gitlab},
	}

	return prompt.Run()
}

func gitAccessToken() (string, error) {
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
