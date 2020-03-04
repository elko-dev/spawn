package builds

// RepoConfigArgs to set repository config
type RepoConfigArgs struct {
	RepoURL string `json:"repo_url"`
}

// RepoConfigResponse message
type RepoConfigResponse struct {
	Message string `json:"message"`
}

// BuildArgs parameters
type BuildArgs struct {
	SourceVersion string `json:"sourceVersion"`
	Debug         bool   `json:"debug"`
}

// BuildResponse object
type BuildResponse struct {
	ID              int    `json:"id"`
	BuildNumber     string `json:"buildNumber"`
	QueueTime       string `json:"queueTime"`
	StartTime       string `json:"startTime"`
	FinishTime      string `json:"finishTime"`
	LastChangedDate string `json:"lastChangedDate"`
	Status          string `json:"status"`
	Result          string `json:"result"`
	SourceBranch    string `json:"sourceBranch"`
	SourceVersion   string `json:"sourceVersion"`
}

// ConfigArgs configuration for application
type ConfigArgs struct {
	Toolsets             Toolsets               `json:"toolsets"`
	EnvironmentVariables []EnvironmentVariables `json:"environmentVariables"`
	Trigger              string                 `json:"trigger"`
	ArtifactVersioning   ArtifactVersioning     `json:"artifactVersioning"`
	BadgeIsEnabled       bool                   `json:"badgeIsEnabled"`
	Signed               bool                   `json:"signed"`
}

type EnvironmentVariables struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Toolsets struct {
	Android      Android      `json:"android"`
	Javascript   Javascript   `json:"javascript"`
	Testcloud    Testcloud    `json:"testcloud"`
	Distribution Distribution `json:"distribution"`
}
type Distribution struct {
	DestinationType string   `json:"destinationType"`
	Destinations    []string `json:"destinations"`
	IsSilent        bool     `json:"isSilent"`
}

type ArtifactVersioning struct {
	BuildNumberFormat string `json:"buildNumberFormat"`
}

type Javascript struct {
	PackageJSONPath    string `json:"packageJsonPath"`
	RunTests           bool   `json:"runTests"`
	ReactNativeVersion string `json:"reactNativeVersion"`
	NodeVersion        string `json:"nodeVersion"`
}

type Android struct {
	GradleWrapperPath string `json:"gradleWrapperPath"`
	Module            string `json:"module"`
	BuildBundle       bool   `json:"buildBundle"`
	BuildVariant      string `json:"buildVariant"`
	RunTests          bool   `json:"runTests"`
	RunLint           bool   `json:"runLint"`
	KeystorePassword  string `json:"keystorePassword"`
	KeyAlias          string `json:"keyAlias"`
	KeyPassword       string `json:"keyPassword"`
	KeystoreFilename  string `json:"keystoreFilename"`
	AutomaticSigning  bool   `json:"automaticSigning"`
}

type Testcloud struct {
	DeviceSelection string `json:"deviceSelection"`
	FrameworkType   string `json:"frameworkType"`
}

type BuildConfigResponse struct {
	Trigger         string `json:"trigger"`
	TestsEnabled    bool   `json:"testsEnabled"`
	BadgeIsEnabled  bool   `json:"badgeIsEnabled"`
	Signed          bool   `json:"signed"`
	CloneFromBranch string `json:"cloneFromBranch"`
	Toolsets        struct {
		Xcode struct {
			ProjectOrWorkspacePath               string `json:"projectOrWorkspacePath"`
			PodfilePath                          string `json:"podfilePath"`
			CartfilePath                         string `json:"cartfilePath"`
			ProvisioningProfileEncoded           string `json:"provisioningProfileEncoded"`
			CertificateEncoded                   string `json:"certificateEncoded"`
			ProvisioningProfileFileID            string `json:"provisioningProfileFileId"`
			CertificateFileID                    string `json:"certificateFileId"`
			ProvisioningProfileUploadID          string `json:"provisioningProfileUploadId"`
			AppExtensionProvisioningProfileFiles []struct {
				FileName               string `json:"fileName"`
				FileID                 string `json:"fileId"`
				UploadID               string `json:"uploadId"`
				TargetBundleIdentifier string `json:"targetBundleIdentifier"`
			} `json:"appExtensionProvisioningProfileFiles"`
			CertificateUploadID         string `json:"certificateUploadId"`
			CertificatePassword         string `json:"certificatePassword"`
			Scheme                      string `json:"scheme"`
			XcodeVersion                string `json:"xcodeVersion"`
			ProvisioningProfileFilename string `json:"provisioningProfileFilename"`
			CertificateFilename         string `json:"certificateFilename"`
			TeamID                      string `json:"teamId"`
			AutomaticSigning            bool   `json:"automaticSigning"`
			XcodeProjectSha             string `json:"xcodeProjectSha"`
			ArchiveConfiguration        string `json:"archiveConfiguration"`
			TargetToArchive             string `json:"targetToArchive"`
			ForceLegacyBuildSystem      bool   `json:"forceLegacyBuildSystem"`
		} `json:"xcode"`
		Javascript struct {
			PackageJSONPath    string `json:"packageJsonPath"`
			RunTests           bool   `json:"runTests"`
			ReactNativeVersion string `json:"reactNativeVersion"`
		} `json:"javascript"`
		Xamarin struct {
			SlnPath       string `json:"slnPath"`
			IsSimBuild    bool   `json:"isSimBuild"`
			Args          string `json:"args"`
			Configuration string `json:"configuration"`
			P12File       string `json:"p12File"`
			P12Pwd        string `json:"p12Pwd"`
			ProvProfile   string `json:"provProfile"`
			MonoVersion   string `json:"monoVersion"`
			SdkBundle     string `json:"sdkBundle"`
			Symlink       string `json:"symlink"`
		} `json:"xamarin"`
		Android struct {
			GradleWrapperPath string `json:"gradleWrapperPath"`
			Module            string `json:"module"`
			BuildVariant      string `json:"buildVariant"`
			RunTests          bool   `json:"runTests"`
			RunLint           bool   `json:"runLint"`
			IsRoot            bool   `json:"isRoot"`
			AutomaticSigning  bool   `json:"automaticSigning"`
			KeystorePassword  string `json:"keystorePassword"`
			KeyAlias          string `json:"keyAlias"`
			KeyPassword       string `json:"keyPassword"`
			KeystoreFilename  string `json:"keystoreFilename"`
			KeystoreEncoded   string `json:"keystoreEncoded"`
		} `json:"android"`
	} `json:"toolsets"`
	ArtifactVersioning struct {
		BuildNumberFormat string `json:"buildNumberFormat"`
	} `json:"artifactVersioning"`
	ID              int `json:"id"`
	AdditionalProp1 struct {
		Branch struct {
			Name   string `json:"name"`
			Commit struct {
				Sha string `json:"sha"`
				URL string `json:"url"`
			} `json:"commit"`
		} `json:"branch"`
		Enabled bool `json:"enabled"`
	} `json:"additionalProp1"`
	AdditionalProp2 struct {
		Branch struct {
			Name   string `json:"name"`
			Commit struct {
				Sha string `json:"sha"`
				URL string `json:"url"`
			} `json:"commit"`
		} `json:"branch"`
		Enabled bool `json:"enabled"`
	} `json:"additionalProp2"`
	AdditionalProp3 struct {
		Branch struct {
			Name   string `json:"name"`
			Commit struct {
				Sha string `json:"sha"`
				URL string `json:"url"`
			} `json:"commit"`
		} `json:"branch"`
		Enabled bool `json:"enabled"`
	} `json:"additionalProp3"`
}

type Keystore struct {
	KeystorePassword string
	KeyAlias         string
	KeyPassword      string
	KeystoreFilename string
}

// CreateConfigArgs default configuration
func CreateConfigArgs(distributionGroupID *string, environmentVariables []EnvironmentVariables, keyStore *Keystore) *ConfigArgs {
	testCloud := Testcloud{DeviceSelection: "top_3_devices", FrameworkType: "Generated"}
	android := Android{
		GradleWrapperPath: "android/gradlew",
		Module:            "app",
		BuildBundle:       false,
		BuildVariant:      "release",
		RunTests:          false,
		RunLint:           true,
		AutomaticSigning:  true,
		KeyAlias:          keyStore.KeyAlias,
		KeyPassword:       keyStore.KeyPassword,
		KeystoreFilename:  keyStore.KeystoreFilename,
		KeystorePassword:  keyStore.KeyPassword,
	}
	javascript := Javascript{
		PackageJSONPath:    "package.json",
		RunTests:           true,
		ReactNativeVersion: "0.61.4",
		NodeVersion:        "10.x",
	}

	distribution := Distribution{
		DestinationType: "groups",
		Destinations:    []string{*distributionGroupID},
		IsSilent:        false,
	}

	toolSets := Toolsets{
		Testcloud:    testCloud,
		Android:      android,
		Javascript:   javascript,
		Distribution: distribution,
	}

	artifactVersioning := ArtifactVersioning{
		BuildNumberFormat: "buildId",
	}

	return &ConfigArgs{
		Toolsets:             toolSets,
		EnvironmentVariables: environmentVariables,
		Trigger:              "continuous",
		ArtifactVersioning:   artifactVersioning,
		BadgeIsEnabled:       false,
		Signed:               true,
	}
}
