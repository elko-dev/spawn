package local

import (
	"os"
	"strings"
	"time"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/file"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Local struct containing logic to interact with Git locally
type Local struct {
}

// Git to interact with git
type Git interface {
	DuplicateRepo(url string, gitToken string, name string, repoURL string) (applications.GitResult, error)
}

// Template interface to replace templated values
type Template interface {
	Replace(name string, path string, fi os.FileInfo) error
}

// DuplicateRepo contains logic to duplicate a repository
func (local Local) DuplicateRepo(url string, gitToken string, name string, repoURL string) (applications.GitResult, error) {

	r, err := git.PlainClone(name, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		println("Clone failed")
		return applications.GitResult{}, err
	}

	template := file.TemplateFile{Name: strings.ToLower(name)}
	err = template.Replace()
	if err != nil {
		println("Template replacement failed")
		return applications.GitResult{}, err
	}

	err = r.DeleteRemote("origin")
	if err != nil {
		println("Delete failed")
		return applications.GitResult{}, err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{repoURL},
	})
	if err != nil {
		println("Create remote failed")
		return applications.GitResult{}, err
	}

	// Adds the new file to the staging area.
	w, err := r.Worktree()
	_, err = w.Add(".")
	if err != nil {
		println("Add failed")
		return applications.GitResult{}, err
	}

	commit, err := w.Commit(name+" configuration", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Spawn",
			Email: "spawn@elko.dev",
			When:  time.Now(),
		},
	})

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: "andrew", // yes, this can be anything except an empty string
			Password: gitToken,
		},
	})
	if err != nil {
		println("Push failed")
		return applications.GitResult{}, err
	}

	return applications.GitResult{RepoURL: repoURL, LatestGitCommit: commit.String()}, err
}

// NewLocal init method
func NewLocal() Local {
	return Local{}
}
