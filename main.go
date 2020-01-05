package main

import (
	"os"

	"gitlab.com/shared-tool-chain/spawn/clip"
)

func main() {
	app := clip.Init(CreateSpawnAction())
	app.Run(os.Args)
}
