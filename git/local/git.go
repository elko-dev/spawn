package local

import (
	"os"
	"strings"
	"time"

	"github.com/elko-dev/spawn/directory"
	"github.com/elko-dev/spawn/file"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Local struct containing logic to interact with Git locally
type Local struct {
}

// GitConfig represents local git configuration
type GitConfig struct {
	RepoURL         string
	LatestGitCommit string
}

// Git to interact with git
type Git interface {
	DuplicateRepo(url string, gitToken string, name string, repoURL string, replacements map[string]string) (GitConfig, error)
}

// Template interface to replace templated values
type Template interface {
	Replace(name string, path string, fi os.FileInfo) error
}

// DuplicateRepo contains logic to duplicate a repository
func (local Local) DuplicateRepo(url string, gitToken string, name string, repoURL string, replacements map[string]string) (GitConfig, error) {

	r, err := git.PlainClone(name, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	if err != nil {
		println("Clone failed")
		return GitConfig{}, err
	}

	template := file.TemplateFile{Name: strings.ToLower(name), Replacements: replacements}
	err = template.Replace()
	if err != nil {
		println("Template file replacement failed")
		return GitConfig{}, err
	}

	dirTemplate := directory.TemplateDirectory{Name: strings.ToLower(name)}
	err = dirTemplate.Replace()
	if err != nil {
		println("Template directory replacement failed")
		return GitConfig{}, err
	}

	err = r.DeleteRemote("origin")
	if err != nil {
		println("Delete failed")
		return GitConfig{}, err
	}

	// Adds the new file to the staging area.
	w, err := r.Worktree()
	//this is a temp hack as changing file names messes up the git tree
	os.Remove(name + "/.git/index")
	err = w.Reset(&git.ResetOptions{Mode: git.SoftReset})

	if err != nil {
		println("Reset failed ", err)
		return GitConfig{}, err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{repoURL},
	})
	if err != nil {
		println("Create remote failed")
		return GitConfig{}, err
	}

	_, err = w.Add(".")
	if err != nil {
		println("Add failed")
		return GitConfig{}, err
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
		return GitConfig{}, err
	}

	return GitConfig{RepoURL: repoURL, LatestGitCommit: commit.String()}, err
}

// NewLocal init method
func NewLocal() Local {
	return Local{}
}
