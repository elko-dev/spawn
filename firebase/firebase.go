package firebase

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/constants"
	"github.com/elko-dev/spawn/gcp"
)

type AppIds struct {
	IosID     string
	AndroidID string
}

// Project to create GCP Project
type Project interface {
	Create(project gcp.ProjectRequest) (gcp.Project, error)
}

// FirebaseProject to add Firebase to GCP project
type FirebaseProject interface {
	Create(gcpProjectID string) (FirebaseProjectResponse, error)
}

// IosApp to create application
type IosApp interface {
	Create(projectID string, request IOSRequest) (IOSResponse, error)
}

// AndroidApp to create application
type AndroidApp interface {
	Create(projectID string, request AndroidRequest) (AndroidResponse, error)
}

// Platform for firebase
type Platform struct {
	projectName     string
	applicationType string
	project         Project
	firebase        FirebaseProject
	iosApp          IosApp
	androidApp      AndroidApp
}

// Create Firebase platform
func (platform Platform) Create() error {
	gcpProjectID := "spawn" + platform.projectName
	gcpProjectName := "spawn" + platform.projectName
	_, err := platform.project.Create(gcp.ProjectRequest{
		//TODO - Add Spawn to project ID to help ensure uniqueness
		ID:   gcpProjectID,
		Name: gcpProjectName,
	})
	if err != nil {
		return err
	}

	firebaseResponse, err := platform.firebase.Create(gcpProjectID)
	if err != nil {
		return err
	}

	_, err = platform.iosApp.Create(firebaseResponse.ID, IOSRequest{
		BundleID:    createBundleID(firebaseResponse.Name),
		DisplayName: createIosName(firebaseResponse.Name),
	})
	if err != nil {
		return err
	}

	_, err = platform.androidApp.Create(firebaseResponse.ID, AndroidRequest{
		BundleID:    createBundleID(firebaseResponse.Name),
		DisplayName: firebaseResponse.Name,
	})
	return err
}

//TODO: this is hardcoded want to move this out
func createBundleID(projectName string) string {
	return "com.elko." + projectName
}

func createIosName(projectName string) string {
	return projectName + "-ios"
}

// GetToken for firebase
func (platform Platform) GetToken() string {
	return ""
}

// GetPlatformType return type of platform (Firebase)
func (platform Platform) GetPlatformType() string {
	return constants.FirebasePlatform
}

// NewPlatform init
func NewPlatform(projectName string,
	applicationType string,
	project Project,
	firebase FirebaseProject,
	iosApp IosApp,
	androidApp AndroidApp) applications.PlatformRepository {
	return Platform{projectName, applicationType, project, firebase, iosApp, androidApp}
}
