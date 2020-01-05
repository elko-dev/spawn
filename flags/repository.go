package flags

import "github.com/urfave/cli"

const (
	projectNameArg = "projectname, pn"
	accessTokenArg = "accesstoken, ac"
	deployTokenArn = "deploytoken, dc"
)

// Repository method to return flags for repository
func Repository() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  projectNameArg,
			Usage: "Project Name `projectname`",
		},
		cli.StringFlag{
			Name:  accessTokenArg,
			Usage: "Access Token `accesstoken`",
		},
		cli.StringFlag{
			Name:  deployTokenArn,
			Usage: "Deploy Token `deploytoken`",
		},
	}
}
