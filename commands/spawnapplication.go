package commands

import (
	"github.com/elko-dev/spawn/applicationtype"
	"github.com/elko-dev/spawn/flags"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ApplicationType interface {
	Create() error
}

// Run is the method to run the CreateRepository command
func Run(factory applicationtype.Factory) cli.Command {
	return cli.Command{
		Name:    "application",
		Aliases: []string{"application"},
		Usage:   "Spawns application",
		Flags:   flags.Repository(),
		Action: func(c *cli.Context) error {
			appType, err := factory.CreateApplicationType()
			if err != nil {
				log.WithFields(log.Fields{}).Error(
					err,
					"\n Spawn encountered an error",
				)
				return err
			}
			return appType.Create()
		},
	}
}
