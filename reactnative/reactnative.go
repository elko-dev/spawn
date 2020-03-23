package reactnative

import "github.com/elko-dev/spawn/applications"

const (
	templateURL                  = "https://github.com/elko-dev/react-native-template.git"
	templateNameReplacement      = "myapp"
	templateIOSIDReplacement     = "1:157175452340:ios:41629cc9b57c97d5f9bde1"
	templateAndroidIDReplacement = "1:157175452340:android:6202e96dce8ed954f9bde1"
	templateWebIDReplacement     = "1:1095344010247:web:6e16064718094a3b8a66a2"
)

type ReactNative struct {
	repo            applications.GitRepo
	ciPlatform      applications.CIPlatform
	mobilePlatform  applications.MobilePlatform
	projectName     string
	includePlatform bool
}

// Create ReactNative Project
func (react ReactNative) Create() error {

	//yes this is a hack to not always include CI
	var mobileApps applications.MobileApps
	var err error

	if react.includePlatform {
		mobileApps, err = react.mobilePlatform.Create()
		if err != nil {
			return err
		}
	}

	// TODO: this fails the interface segregration principle.  Need to refactor
	response, err := react.repo.CreateGitRepository(
		react.projectName,
		templateURL,
		"",
		createReplacements(react.projectName, mobileApps.IOS.ID, mobileApps.Android.ID, mobileApps.Web.AppID, react.includePlatform))

	if err != nil {
		return err
	}

	return react.ciPlatform.Create(response.RepoURL, response.RepoID, response.LatestGitCommit, react.repo.GetRepoType())
}

func createReplacements(projectName string, iosID string, androidID string, webID string, includePlatform bool) map[string]string {
	replacements := make(map[string]string)
	replacements[templateNameReplacement] = projectName

	if includePlatform {
		replacements[templateIOSIDReplacement] = iosID
		replacements[templateAndroidIDReplacement] = androidID
		replacements[templateWebIDReplacement] = webID
	}
	return replacements
}

// NewReactNative init function
func NewReactNative(repo applications.GitRepo,
	ciPlatform applications.CIPlatform,
	mobilePlatform applications.MobilePlatform,
	projectName string,
	includePlatform bool) ReactNative {
	return ReactNative{repo, ciPlatform, mobilePlatform, projectName, includePlatform}
}
