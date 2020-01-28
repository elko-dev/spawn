package selections

import (
	"github.com/elko-dev/spawn/constants"
	"github.com/manifoldco/promptui"
)

// ApplicationType prompts user to select type of application
func ApplicationType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Application Type",
		Items: []string{constants.NodeServerType, constants.ReactClientLanguageType},
	}

	return prompt.Run()
}
