package builds

import (
	"reflect"
	"testing"
)

const androidKeyStoreKey = "ANDROID_KEYSTORE_KEY"

func TestCreateConfigArgs(t *testing.T) {
	encryptToken := "asdfasdf"

	environmentVariables := []EnvironmentVariables{
		EnvironmentVariables{
			Name:  androidKeyStoreKey,
			Value: encryptToken,
		}}
	distributionID := "distID"
	keyStore := Keystore{}
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
		Destinations:    []string{distributionID},
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

	type args struct {
		distributionGroupID  *string
		environmentVariables []EnvironmentVariables
		keyStore             *Keystore
	}
	tests := []struct {
		name string
		args args
		want *ConfigArgs
	}{
		struct {
			name string
			args args
			want *ConfigArgs
		}{
			name: "KeyStore Is Added As ENV",
			args: args{&distributionID, environmentVariables, &keyStore},
			want: &ConfigArgs{
				Toolsets:             toolSets,
				EnvironmentVariables: environmentVariables,
				Trigger:              "continuous",
				ArtifactVersioning:   artifactVersioning,
				BadgeIsEnabled:       false,
				Signed:               true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConfigArgs(tt.args.distributionGroupID, tt.args.environmentVariables, tt.args.keyStore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConfigArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
