package nodejs

import (
	"testing"

	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	gomock "github.com/golang/mock/gomock"
)

const (
	platformName    = "somename"
	applicationType = "type"
	framework       = constants.GraphQLHerokuPlatform
)

func TestRetrievesProjectNameAndCreatesPlatformWithName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gitFactoryMock := applications.NewMockGitFactory(ctrl)
	platformFactoryMock := applications.NewMockPlatformFactory(ctrl)
	platformMock := NewMockPrompt(ctrl)

	gitFactoryMock.EXPECT().Create(platformName).Return(nil, nil)
	platformFactoryMock.EXPECT().Create(platformName, applicationType).Return(nil, nil)
	platformMock.EXPECT().forAppName().Return(platformName, nil)
	platformMock.EXPECT().forFramework().Return(framework, nil)
	factory := NewFactory(gitFactoryMock, platformFactoryMock, platformMock)

	project, _ := factory.Create(applicationType)

	if project == nil {
		t.Log("project not returned")
		t.Fail()
	}

}
