package reactnative

import "github.com/elko-dev/spawn/applications"

const (
	templateURL                  = "https://github.com/elko-dev/react-native-template.git"
	templateNameReplacement      = "myapp"
	templateIOSIDReplacement     = "1:157175452340:ios:41629cc9b57c97d5f9bde1"
	templateAndroidIDReplacement = "1:157175452340:android:6202e96dce8ed954f9bde1"
)

type ReactNative struct {
	repo           applications.GitRepo
	ciPlatform     applications.CIPlatform
	mobilePlatform applications.MobilePlatform
	projectName    string
}

// Create ReactNative Project
func (react ReactNative) Create() error {

	mobileApps, err := react.mobilePlatform.Create()

	if err != nil {
		return err
	}

	// TODO: this fails the interface segregration principle.  Need to refactor
	response, err := react.repo.CreateGitRepository(
		react.projectName,
		templateURL,
		"",
		createReplacements(react.projectName, mobileApps.IOS.ID, mobileApps.Android.ID))

	if err != nil {
		return err
	}

	return react.ciPlatform.Create(response.RepoURL, response.LatestGitCommit)
}

func createReplacements(projectName string, iosID string, androidID string) map[string]string {
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName
	replacements[templateIOSIDReplacement] = iosID
	replacements[templateAndroidIDReplacement] = androidID
	return replacements
}

// NewReactNative init function
func NewReactNative(repo applications.GitRepo, ciPlatform applications.CIPlatform, mobilePlatform applications.MobilePlatform, projectName string) ReactNative {
	return ReactNative{repo, ciPlatform, mobilePlatform, projectName}
}
