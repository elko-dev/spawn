package applications

// React struct to create Node aplication
type React struct {
	name        string
	accessToken string
	deployToken string
	teamName    string
	Repo        GitRepository
	Platform    PlatformRepository
}

// Create is a function to generate a NodeJS application
func (react React) Create(environment string) error {

	return nil
}

// NewReact init function
func NewReact(gitRepository GitRepository, platform PlatformRepository) React {
	return React{Repo: gitRepository, Platform: platform}
}
