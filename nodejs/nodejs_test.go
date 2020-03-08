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
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName
	mockGitRepo := applications.NewMockGitRepo(ctrl)
	mockPlatform := applications.NewMockPlatformRepository(ctrl)

	mockPlatform.EXPECT().Create().Return(nil)
	mockPlatform.EXPECT().GetToken().Return(token)
	mockGitRepo.EXPECT().CreateGitRepository(projectName, graphQLHerokuTemplateURL, token, replacements).Return(applications.GitResult{}, nil)

	node := NewNode(mockGitRepo, mockPlatform, projectName, framework)

	err := node.Create()

	if err != nil {
		t.Log("error encountered when non expected ", err)
		t.Fail()
		return
	}

}

func TestReturnsErrorWhenBasedFrameworkIsProvided(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGitRepo := applications.NewMockGitRepo(ctrl)
	mockPlatform := applications.NewMockPlatformRepository(ctrl)

	node := NewNode(mockGitRepo, mockPlatform, projectName, "DOESNOTEXISTFRAMEWORK")

	err := node.Create()

	if err == nil {
		t.Log("No error encountered when one should have been returned")
		t.Fail()
		return
	}

}

func TestHerokuTemplateIsProvidedWhenPlatformIsHeroku(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName

	mockGitRepo := applications.NewMockGitRepo(ctrl)
	mockPlatform := applications.NewMockPlatformRepository(ctrl)

	mockPlatform.EXPECT().Create().Return(nil)
	mockPlatform.EXPECT().GetToken().Return(token)
	mockGitRepo.EXPECT().CreateGitRepository(projectName, graphQLHerokuTemplateURL, token, replacements).Return(applications.GitResult{}, nil)

	node := NewNode(mockGitRepo, mockPlatform, projectName, framework)

	err := node.Create()

	if err != nil {
		t.Log("error encountered when non expected ", err)
		t.Fail()
		return
	}

}

func Test_getTemplateURL(t *testing.T) {
	type args struct {
		platformType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{functionsTemplateURL, args{platformType: constants.AzureFunctions}, functionsTemplateURL, false},
		struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{expressHerokuTemplateURL, args{platformType: constants.ExpressHerokuPlatform}, expressHerokuTemplateURL, false},
		struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{graphQLHerokuTemplateURL, args{platformType: constants.GraphQLHerokuPlatform}, graphQLHerokuTemplateURL, false},
		struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{"Invalid type", args{platformType: constants.HerokuPlatform}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTemplateURL(tt.args.platformType)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTemplateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
