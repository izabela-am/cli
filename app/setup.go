package app

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Setup() *cli.App {
	toolkit := cli.NewApp()

	toolkit.Name = "Toolkit"
	toolkit.Usage = "A set of scripts I use a lot"
	toolkit.UsageText = "Write usage here"

	toolkit.Commands = []*cli.Command{
		{
			Name:    "Network Device Scanner",
			Aliases: []string{"ncd"},
			Usage:   "Lists all devices connected in your network",
			Action: func(context *cli.Context) error {
				ConnectedDevices()
				return nil
			},
		},
		{
			Name:    "Password Leak Checker",
			Aliases: []string{"plc"},
			Usage:   "Checks a given password against a database of leaked credentials",
			Action: func(context *cli.Context) error {
				PasswordLeaks()
				return nil
			},
		},
	}

	err := toolkit.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	return toolkit
}
