package validations

import (
	"errors"
	"strings"
)

// YOrNValidation validates whether input is y or n
func YOrNValidation(input string) error {
	lowercaseInput := strings.ToLower(input)
	if lowercaseInput != "y" && lowercaseInput != "n" {
		return errors.New("must select y or n")
	}
	return nil
}

// GitValidation validates whether url is a .git url
func GitValidation(input string) error {
	if !strings.HasSuffix(input, ".git") {
		return errors.New("template must be a .git url")
	}

	return nil
}
