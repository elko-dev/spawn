//+build wireinject

package main //The above newline is required by wire https://github.com/google/wire/issues/117

import (
	"github.com/google/wire"
	"github.com/elko-dev/spawn/actions"
)

// CreateSpawnAction dependency
func CreateSpawnAction() actions.SpawnAction {
	panic(wire.Build(
		actions.NewSpawnAction,
	))
}

// func CreateGitlabRepistory() applications.GitRepository {
// 	panic(wire.Build(
// 		git.NewLocal,
// 		git.NewGitlabRepository,
// 	))
// }
