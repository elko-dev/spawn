package prompt

// Command interface containing all prompted commands
type Command interface {
	ApplicationType() (string, error)
	ServerType() (string, error)
	ClientLanguageType(applicationType string) (string, error)
	Platform() (string, error)
}

// Selection is a struct containing functionality to determine user application to spawn
type Selection struct {
	Command Command
}

// UserSelections represents the responses from users
type UserSelections struct {
	ApplicationType    string
	ServerType         string
	ClientLanguageType string
}

// Application guides user to select an application to Spawn
func (selection Selection) Application() (UserSelections, error) {
	// Selection process:
	// Application type: Web or Mobile
	// Server Language: NodeJS
	// Client Language (Mobile): React Native, Flutter, etc
	// Client Language (Web): React
	// Platform: Heroku
	// Version Control: Gitlab
	// CI/CD: GitlabCI
	applicationType, err := selection.Command.ApplicationType()

	if err != nil {
		return UserSelections{}, err
	}

	serverType, err := selection.Command.ServerType()

	if err != nil {
		return UserSelections{}, err
	}

	clientLanguageType, err := selection.Command.ClientLanguageType(applicationType)

	if err != nil {
		return UserSelections{}, err

	}

	return UserSelections{
		ApplicationType:    applicationType,
		ServerType:         serverType,
		ClientLanguageType: clientLanguageType,
	}, nil
}
