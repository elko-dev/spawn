package main

import (
	"os"

	"github.com/elko-dev/spawn/clip"
)

func main() {

	applicationFactory := CreateFactory()
	app := clip.Init(applicationFactory)
	app.Run(os.Args)
}
