package nodejs

import (
	"testing"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	gomock "github.com/golang/mock/gomock"
)

const (
	projectName = "somename"
	token       = "token"
)

func TestFunctionsTemplateIsProvidedWhenPlatformIsFunctions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGitRepo := applications.NewMockGitRepo(ctrl)
	mockPlatform := applications.NewMockPlatformRepository(ctrl)

	mockPlatform.EXPECT().Create().Return(nil)
	mockPlatform.EXPECT().GetToken().Return(token)
	mockPlatform.EXPECT().GetPlatformType().Return(constants.AzureFunctions)
	mockGitRepo.EXPECT().CreateGitRepository(projectName, functionsTemplateURL, token).Return(nil)

	node := NewNode(mockGitRepo, mockPlatform, projectName)

	err := node.Create()

	if err != nil {
		t.Log("error encountered when non expected ", err)
		t.Fail()
		return
	}

}

func TestHerokuTemplateIsProvidedWhenPlatformIsHeroku(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGitRepo := applications.NewMockGitRepo(ctrl)
	mockPlatform := applications.NewMockPlatformRepository(ctrl)

	mockPlatform.EXPECT().Create().Return(nil)
	mockPlatform.EXPECT().GetToken().Return(token)
	mockPlatform.EXPECT().GetPlatformType().Return(constants.HerokuPlatform)
	mockGitRepo.EXPECT().CreateGitRepository(projectName, herokuTemplateURL, token).Return(nil)

	node := NewNode(mockGitRepo, mockPlatform, projectName)

	err := node.Create()

	if err != nil {
		t.Log("error encountered when non expected ", err)
		t.Fail()
		return
	}

}
