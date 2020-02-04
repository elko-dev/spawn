package gitlab

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forGroupId() (string, error){
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
		Label:    "Gitlab Access Token",
		Validate: accessTokenValidate,
		Mask:     '*',
	}

	return accessTokenPrompt.Run()
}

// NewPrompts init func
func NewPrompts() Prompts {
	return Prompts{}
}
