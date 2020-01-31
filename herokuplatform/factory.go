package herokuplatform

// Factory to create a platform
type Factory struct {
	prompt Prompt
}

// Prompt interface to retrieve values from
type Prompt interface {
	forEnvironments() ([]string, error)
	forHerokuTeamName() (string, error)
	forPlatformToken() (string, error)
}

// Create a platform repo
func (factory Factory) Create(projectName string) (Heroku, error) {
	// fields to create
	// applicationType  string
	envs, err := factory.prompt.forEnvironments()

	if err != nil {
		return Heroku{}, err
	}

	teamName, err := factory.prompt.forHerokuTeamName()
	if err != nil {
		return Heroku{}, err
	}

	token, err := factory.prompt.forPlatformToken()
	if err != nil {
		return Heroku{}, err
	}

	return NewHeroku(token, envs, projectName, teamName, ""), nil
}

// NewFactory init
func NewFactory(prompt Prompt) Factory {
	return Factory{prompt}
}
