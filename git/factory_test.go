package git

import (
	"testing"

	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/git/ados"
	"github.com/elko-dev/spawn/git/gitlab"
	gomock "github.com/golang/mock/gomock"
)

const (
	projectName = "name"
)

func TestWhenGitlabIsSelectedGitlabRepoIsCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrompt := NewMockPrompt(ctrl)
	mockPrompt.EXPECT().forGitRepository().Return(constants.Gitlab, nil)
	factory := NewFactory(mockPrompt)
	repo, _ := factory.Create(projectName)
	if !isGitlabType(repo) {
		t.Log("Expected gitlab type returned")
		t.Fail()
		return
	}

}
func TestWhenADOSIsSelectedADOSRepoIsCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPrompt := NewMockPrompt(ctrl)
	mockPrompt.EXPECT().forGitRepository().Return(constants.ADOS, nil)
	factory := NewFactory(mockPrompt)
	repo, _ := factory.Create(projectName)
	if !isADOSType(repo) {
		t.Log("Expected ados type returned")
		t.Fail()
		return
	}

}

func isGitlabType(t interface{}) bool {
	switch t.(type) {

	case gitlab.GitlabRepo:
		return true
	default:
		return false
	}

}

func isADOSType(t interface{}) bool {
	switch t.(type) {

	case ados.Repository:
		return true
	default:
		return false
	}

}
