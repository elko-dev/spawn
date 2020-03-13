package firebase

import (
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/gcp"
)

// Factory to create firebase project
type Factory struct {
}

// Create Firebase factory
func (factory Factory) Create(projectName string, applicationType string) (applications.MobilePlatform, error) {
	project := gcp.NewProjectClient()
	firebaseProject := NewProjectClient()

	iosApp := NewIOSClient()
	androidApp := NewAndroidClient()
	webApp := NewWebClient()
	return NewPlatform(projectName, applicationType, project, firebaseProject, iosApp, androidApp, webApp), nil
}

// NewFactory init
func NewFactory() applications.MobilePlatformFactory {
	return Factory{}
}
