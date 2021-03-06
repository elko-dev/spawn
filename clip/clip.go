// Package clip is the CLI interface for spawn
package clip

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/commands"
	"github.com/urfave/cli"
)

// Command is an interface defining structure for CLI commands
type Command interface {
	Run(input *CommandArgs) error
}

// CommandArgs is a struct representing params required for command
type CommandArgs struct {
	ProjectName string
}

// Init ... this is a basic fn
func Init(factory applicationtype.Factory) *cli.App {

	app := cli.NewApp()
	app.Name = "spawn"
	app.Usage = "Spawn creates project scaffolding, integrating GitLab with Heroku."
	app.Commands = []cli.Command{
		commands.Run(factory),
	}
	return app
}
