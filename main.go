package main

import (
	"os"

	"github.com/elko-dev/spawn/clip"
	log "github.com/sirupsen/logrus"
)

func main() {

	//TODO: Factor out to set log levels via cli

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	applicationFactory := CreateApplicationTypeFacotry()
	app := clip.Init(applicationFactory)
	app.Run(os.Args)
}
