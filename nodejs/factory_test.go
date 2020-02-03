package nodejs

import (
	"testing"

	"github.com/elko-dev/spawn/applications"
	gomock "github.com/golang/mock/gomock"
)

const (
	platformName = "somename"
)

func TestRetrievesProjectNameAndCreatesPlatformWithName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gitFactoryMock := applications.NewMockGitFactory(ctrl)
	platformFactoryMock := applications.NewMockPlatformFactory(ctrl)
	platformMock := NewMockPrompt(ctrl)

	gitFactoryMock.EXPECT().Create(platformName).Return(nil, nil)
	platformFactoryMock.EXPECT().Create(platformName).Return(nil, nil)
	platformMock.EXPECT().forAppName().Return(platformName, nil)

	factory := NewFactory(gitFactoryMock, platformFactoryMock, platformMock)

	project, _ := factory.Create()

	if project == nil {
		t.Log("project not returned")
		t.Fail()
	}

}
