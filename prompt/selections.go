package prompt

// Command interface containing all prompted commands
type Command interface {
	ProjectName() (string, error)
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
	ProjectName        string
	ApplicationType    string
	ServerType         string
	ClientLanguageType string
	Platform           string
	VersionControl     string
	CIServer           string
}

// Application guides user to select an application to Spawn
func (selection Selection) Application() (UserSelections, error) {

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

	platform, err := selection.Command.Platform()

	if err != nil {
		return UserSelections{}, err

	}

	projectName, err := selection.Command.ProjectName()

	if err != nil {
		return UserSelections{}, err

	}

	return UserSelections{
		ApplicationType:    applicationType,
		ServerType:         serverType,
		ClientLanguageType: clientLanguageType,
		Platform:           platform,
		ProjectName:        projectName,
		VersionControl:     "Gitlab",
		CIServer:           "GitlabCI",
	}, nil
}
