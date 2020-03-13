package firebase

import (
	"github.com/elko-dev/spawn/applications"
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
	Create(projectID string, request IOSRequest) (applications.IOSApp, error)
}

// WebApp to create application
type WebApp interface {
	Create(projectID string, request WebRequest) (applications.WebApp, error)
}

// AndroidApp to create application
type AndroidApp interface {
	Create(projectID string, request AndroidRequest) (applications.AndroidApp, error)
}

// Platform for firebase
type Platform struct {
	projectName     string
	applicationType string
	project         Project
	firebase        FirebaseProject
	iosApp          IosApp
	androidApp      AndroidApp
	webApp          WebApp
}

// Create Firebase platform
func (platform Platform) Create() (applications.MobileApps, error) {
	gcpProjectID := "spawn" + platform.projectName
	gcpProjectName := "spawn" + platform.projectName
	_, err := platform.project.Create(gcp.ProjectRequest{
		//TODO - Add Spawn to project ID to help ensure uniqueness
		ID:   gcpProjectID,
		Name: gcpProjectName,
	})
	if err != nil {
		return applications.MobileApps{}, err
	}

	firebaseResponse, err := platform.firebase.Create(gcpProjectID)
	if err != nil {
		return applications.MobileApps{}, err
	}

	iosApp, err := platform.iosApp.Create(firebaseResponse.ID, IOSRequest{
		BundleID:    createBundleID(firebaseResponse.Name),
		DisplayName: createIosName(firebaseResponse.Name),
	})
	if err != nil {
		return applications.MobileApps{}, err
	}

	androidApp, err := platform.androidApp.Create(firebaseResponse.ID, AndroidRequest{
		BundleID:    createBundleID(firebaseResponse.Name),
		DisplayName: createAndroidName(firebaseResponse.Name),
	})

	webApp, err := platform.webApp.Create(firebaseResponse.ID, WebRequest{
		DisplayName: createWebName(firebaseResponse.Name),
	})
	return applications.MobileApps{
		IOS:     iosApp,
		Android: androidApp,
		Web:     webApp,
	}, err
}

//TODO: this is hardcoded want to move this out
func createBundleID(projectName string) string {
	return "com.elko." + projectName
}

func createAndroidName(projectName string) string {
	return projectName + "-android"
}
func createWebName(projectName string) string {
	return projectName + "-web"
}

func createIosName(projectName string) string {
	return projectName + "-ios"
}

// NewPlatform init
func NewPlatform(projectName string,
	applicationType string,
	project Project,
	firebase FirebaseProject,
	iosApp IosApp,
	androidApp AndroidApp,
	webApp WebApp) applications.MobilePlatform {
	return Platform{projectName, applicationType, project, firebase, iosApp, androidApp, webApp}
}
