package main

import (
	"os"

	"github.com/elko-dev/spawn/clip"
)

func main() {
	app := clip.Init(CreateSpawnAction())
	app.Run(os.Args)
}
