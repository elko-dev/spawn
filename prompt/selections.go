package prompt

// Command interface containing all prompted commands
type Command interface {
	ProjectName() (string, error)
	ApplicationType() (string, error)
	ServerType() (string, error)
	ClientLanguageType(applicationType string) (string, error)
}

// PlatformCommand interface containing all prompted commands for Platform
type PlatformCommand interface {
	Platform() (string, string, error)
}

// GitCommand to retrieve git values
type GitCommand interface {
	Token() (string, error)
	Repository() (string, error)
}

// PlatformInput struct to retrieve user commands
type PlatformInput struct {
	Token    string
	TeamName string
}

// Selection is a struct containing functionality to determine user application to spawn
type Selection struct {
	Command         Command
	PlatformCommand PlatformCommand
	GitCommand      GitCommand
}

// UserSelections represents the responses from users
type UserSelections struct {
	ProjectName        string
	ApplicationType    string
	ServerType         string
	ClientLanguageType string
	PlatformToken      string
	PlatformTeamName   string
	Platform           string
	GitToken           string
	VersionControl     string
	CIServer           string
}

// Application guides user to select an application to Spawn
func (selection Selection) Application() (UserSelections, error) {

	applicationType, err := selection.Command.ApplicationType()

	if err != nil {
		return UserSelections{}, err
	}

	gitRepository, err := selection.GitCommand.Repository()

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

	token, teamName, err := selection.PlatformCommand.Platform()

	if err != nil {
		return UserSelections{}, err
	}

	projectName, err := selection.Command.ProjectName()

	if err != nil {
		return UserSelections{}, err
	}

	gitToken, err := selection.GitCommand.Token()

	if err != nil {
		return UserSelections{}, err
	}

	return UserSelections{
		ApplicationType:    applicationType,
		ServerType:         serverType,
		ClientLanguageType: clientLanguageType,
		PlatformToken:      token,
		PlatformTeamName:   teamName,
		//TODO: check for supported selections when other options are allowed
		Platform:       applicationType,
		ProjectName:    projectName,
		GitToken:       gitToken,
		VersionControl: gitRepository,
		CIServer:       "GitlabCI",
	}, nil
}
