//+build wireinject

package main //The above newline is required by wire https://github.com/google/wire/issues/117

import (
	"github.com/google/wire"
	"gitlab.com/shared-tool-chain/spawn/actions"
	"gitlab.com/shared-tool-chain/spawn/git"
	"gitlab.com/shared-tool-chain/spawn/platform"
)

// CreateSpawnAction dependency
func CreateSpawnAction() actions.SpawnAction {
	panic(wire.Build(
		git.NewGitlabRepository,
		platform.NewHerokuPlatform,
		actions.NewSpawnAction,
	))
}
