package git

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/git/ados"
	"github.com/elko-dev/spawn/git/gitlab"
)

// Factory to create GitRepo
type Factory struct {
	prompt Prompt
}

// Prompt user for repo details
type Prompt interface {
	forGitRepository() (string, error)
}

// Create git repo
func (factory Factory) Create(projectName string) (applications.GitRepo, error) {

	//select repo
	repoType, err := factory.prompt.forGitRepository()

	if err != nil {
		return nil, err
	}

	if repoType == constants.ADOS {
		return ados.NewRepository(ados.NewPrompts()), nil
	}

	//construct
	return gitlab.NewGitlabRepo(gitlab.NewPrompts()), nil
}

// NewFactory init
func NewFactory(prompt Prompt) applications.GitFactory {
	return Factory{prompt}
}
