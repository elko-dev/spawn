package main

import (
	"os"

	"github.com/elko-dev/spawn/clip"
)

func main() {
	app := clip.Init()
	app.Run(os.Args)
}
