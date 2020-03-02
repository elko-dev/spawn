//+build wireinject

package main //The above newline is required by wire https://github.com/google/wire/issues/117

import (
	"github.com/elko-dev/spawn/api"
	"github.com/elko-dev/spawn/appcenter"
	"github.com/elko-dev/spawn/applications"
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/azurefunctions"
	"github.com/elko-dev/spawn/functions"
	"github.com/elko-dev/spawn/git"
	"github.com/elko-dev/spawn/herokuplatform"
	"github.com/elko-dev/spawn/mobile"
	"github.com/elko-dev/spawn/nodejs"
	"github.com/elko-dev/spawn/platform"
	"github.com/elko-dev/spawn/react"
	"github.com/elko-dev/spawn/reactnative"
	"github.com/elko-dev/spawn/web"
	"github.com/google/wire"
)

func CreateApplicationTypeFacotry() applicationtype.Factory {
	panic(wire.Build(
		applicationtype.NewPrompts,
		CreateWebFactory,
		CreateFunctionsTypeFactory,
		CreateMobileTypeFactory,
		CreateAPITypeFactory,
		applicationtype.NewFactory,
	))
}

func CreateAPITypeFactory() applicationtype.APITypeFactory {
	panic(wire.Build(
		CreateNodeJsFactory,
		web.NewPrompts,
		api.NewFactory,
	))

}
func CreateMobileTypeFactory() applicationtype.MobileTypeFactory {
	panic(wire.Build(
		CreateNodeJsFactory,
		CreateReactNativeFactory,
		web.NewPrompts,
		mobile.NewFactory,
	))
}
func CreateFunctionsTypeFactory() applicationtype.FunctionTypeFactory {
	panic(wire.Build(
		CreateNodeJsFactory,
		functions.NewFactory,
	))
}

func CreateWebFactory() applicationtype.WebTypeFactory {
	panic(wire.Build(
		CreateNodeJsFactory,
		CreateReactFactory,
		web.NewPrompts,
		web.NewWebFactory,
	))
}

func CreateNodeJsFactory() web.ServerAppFactory {
	panic(wire.Build(
		CreateGitFactory,
		CreatePlatformFactory,
		nodejs.NewPrompts,
		nodejs.NewFactory,
	))
}
func CreateReactFactory() web.ClientAppFactory {
	panic(wire.Build(
		CreateGitFactory,
		CreatePlatformFactory,
		react.NewPrompts,
		react.NewFactory,
	))
}
func CreateReactNativeFactory() web.ClientAppFactory {
	panic(wire.Build(
		CreateGitFactory,
		CreateAppCenterFactory,
		reactnative.NewPrompts,
		reactnative.NewFactory,
	))
}

func CreateAppCenterFactory() applications.CIFactory {
	panic(wire.Build(
		appcenter.NewPrompts,
		appcenter.NewFactory,
	))
}
func CreateGitFactory() applications.GitFactory {
	panic(wire.Build(
		git.NewPrompts,
		git.NewFactory,
	))
}

func CreatePlatformFactory() applications.PlatformFactory {
	panic(wire.Build(
		platform.NewPrompts,
		CreateHerokuFactory,
		CreateAzureFunctionsFactory,
		platform.NewFactory,
	))
}

func CreateAzureFunctionsFactory() platform.FunctionsPlatformFactory {
	panic(wire.Build(
		azurefunctions.NewPrompts,
		azurefunctions.NewFactory,
	))
}

func CreateHerokuFactory() platform.HerokuPlatformFactory {
	panic(wire.Build(
		herokuplatform.NewPrompts,
		herokuplatform.NewFactory,
	))
}
