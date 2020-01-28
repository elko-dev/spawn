package prompt

import (
	"github.com/elko-dev/spawn/constants"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/manifoldco/promptui"
)

// UserCommands struct to retrieve user commands
type UserCommands struct {
}

// Selection process:
// Application type: Web or Mobile
// Server Language: NodeJS
// Client Language (Mobile): React Native, Flutter, etc
// Client Language (Web): React
// Platform: Heroku
// Version Control: Gitlab
// CI/CD: GitlabCI

// ProjectName func
func (userCommands UserCommands) ProjectName() (string, error) {
	selection, err := selectProjectName()
	return selection, err
}

// ApplicationType func
func (userCommands UserCommands) ApplicationType() (string, error) {
	_, selection, err := selectApplicationType()
	return selection, err
}

// ServerType func
func (userCommands UserCommands) ServerType() (string, error) {
	_, selection, err := selectServerType()
	return selection, err
}

// ClientLanguageType func
func (userCommands UserCommands) ClientLanguageType(applicationType string) (string, error) {
	clientLanguageType := getClientLangaugeSelections(applicationType)
	_, selection, err := selectClientLanguageTypes(clientLanguageType)
	return selection, err
}

// Platform func
func (userCommands UserCommands) Platform() (string, error) {
	return "Heroku", nil
}

func getClientLangaugeSelections(applicationType string) []string {
	if applicationType == "Web" {
		return []string{constants.ReactClientLanguageType}
	}
	if applicationType == "Mobile" {
		return []string{"React Native"}
	}
	return make([]string, 0, 0)
}

func selectClientLanguageTypes(selections []string) (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Client Language",
		Items: selections,
	}

	return prompt.Run()
}

func selectServerType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Backend Language",
		Items: []string{constants.NodeServerType},
	}

	return prompt.Run()
}

func selectApplicationType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Application Type",
		Items: []string{"Web", "Mobile"},
	}

	return prompt.Run()
}
func selectProjectName() (string, error) {
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

// PlatformToken prompts user for heroku token
func PlatformToken() (string, error) {
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

