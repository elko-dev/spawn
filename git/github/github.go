package github

import (
	"context"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/git/local"
	"github.com/google/go-github/github"

	"golang.org/x/oauth2"

	log "github.com/sirupsen/logrus"
)

type GithubRepo struct {
	Git    local.Git
	prompt Prompt
}

type Prompt interface {
	forGitToken() (string, error)
}

// CreateGitRepository to create repo
func (git GithubRepo) CreateGitRepository(repositoryName string, templateURL string, platformToken string, replacements map[string]string) (applications.GitResult, error) {
	logContext := log.WithFields(log.Fields{
		"repositoryName": repositoryName,
		"templateURL":    templateURL,
	})
	gitToken, err := git.prompt.forGitToken()
	logContext.Debug("About to create Github application")

	if err != nil {
		logContext.Error("Error retrieving git token")
		return applications.GitResult{}, err
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: gitToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	private := false
	r := &github.Repository{Name: &repositoryName, Private: &private, Description: &repositoryName}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		logContext.Error("Error creating Github repository")
		return applications.GitResult{}, err
	}

	return git.Git.DuplicateRepo(templateURL, gitToken, repo.GetName(), *repo.CloneURL, replacements)
}

// NewGithubRepo init
func NewGithubRepo(prompt Prompt) applications.GitRepo {
	return GithubRepo{local.NewLocal(), prompt}
}
