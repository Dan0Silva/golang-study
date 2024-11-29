package app

import (
	"github.com/urfave/cli"
)

// return a cli application
func CreateApp() *cli.App {
	app := cli.NewApp()

	app.Name = "cli application example"
	app.Usage = "search by IP and name from servers on internet"

	defaultFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search by ip",
			Flags:  defaultFlags,
			Action: getIPs,
		},
		{
			Name:   "servers",
			Usage:  "search by server",
			Flags:  defaultFlags,
			Action: getServers,
		},
	}

	return app
}
