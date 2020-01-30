package prompt

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

// GitlabCommand specifies methods to retrieve git configuration
type GitlabCommand struct {
}

// Token allows user to input token
func (gitlab GitlabCommand) Token() (string, error) {
	token, err := gitlabAccessToken()

	if err != nil {
		return "", err
	}
	return token, nil
}

func gitlabAccessToken() (string, error) {
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
