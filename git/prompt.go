package git

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/manifoldco/promptui"
)

type Prompts struct {
}

func (prompt Prompts) forGitRepository() (string, error) {
	_, selection, err := selectGitRepository()
	return selection, err
}

func selectGitRepository() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Version Control",
		Items: []string{constants.Gitlab, constants.ADOS, constants.Github},
	}

	return prompt.Run()
}

func NewPrompts() Prompt {
	return Prompts{}
}
