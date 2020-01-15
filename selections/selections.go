package selections

import (
	"github.com/manifoldco/promptui"
)

// ApplicationType prompts user to select type of application
func ApplicationType() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select Application Type",
		Items: []string{"NodeJs", "React"},
	}

	return prompt.Run()
}
