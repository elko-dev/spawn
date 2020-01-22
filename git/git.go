package git

import (
	"os"
	"strings"
	"time"

	"github.com/elko-dev/spawn/file"
	"github.com/elko-dev/spawn/git/api"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Local struct containing logic to interact with Git locally
type Local struct {
}

// Template interface to replace templated values
type Template interface {
	Replace(name string, path string, fi os.FileInfo) error
}

// DuplicateRepo contains logic to duplicate a repository
func (local Local) DuplicateRepo(url string, gitToken string, repository api.GitRepository) error {

	r, err := git.PlainClone(repository.Name, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		println("Clone failed")
		println(err.Error())
		return err
	}

	template := file.TemplateFile{Name: strings.ToLower(repository.Name)}
	err = template.Replace()
	if err != nil {
		println("Template replacement failed")
		println(err.Error())
		return err
	}

	err = r.DeleteRemote("origin")
	if err != nil {
		println("Delete failed")
		println(err.Error())
		return err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{repository.URL},
	})
	if err != nil {
		println("Create remote failed")
		println(err.Error())
		return err
	}

	// Adds the new file to the staging area.
	w, err := r.Worktree()
	_, err = w.Add(".")
	if err != nil {
		println("Add failed")
		println(err.Error())
		return err
	}

	_, err = w.Commit(repository.Name+" configuration", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Spawn",
			Email: "spawn@elko.dev",
			When:  time.Now(),
		},
	})

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: gitToken,
		},
	})
	if err != nil {
		println("Push failed")
		println(err.Error())
		return err
	}

	return nil
}

// NewLocal init method
func NewLocal() Git {
	return Local{}
}
