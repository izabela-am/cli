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
			Name:    "html_parser",
			Aliases: []string{"hp"},
			Usage:   "TODO",
			Action: func(context *cli.Context) error {
				HTMLParser(context.Args().First())
				return nil
			},
		},
		{
			Name:    "connected_devices",
			Aliases: []string{"ncd"},
			Usage:   "TODO",
			Action: func(context *cli.Context) error {
				ConnectedDevices()
				return nil
			},
		},
		{
			Name:    "Linux Dependency Manager",
			Aliases: []string{"ldm"},
			Usage:   "TODO",
			Action: func(context *cli.Context) error {
				// TODO
				return nil
			},
		},
		{
			Name:    "FTP Bruteforcer",
			Aliases: []string{"ftpb"},
			Usage:   "TODO",
			Action: func(context *cli.Context) error {
				// TODO
				return nil
			},
		},
		{
			Name:    "Portscan",
			Aliases: []string{"ps"},
			Usage:   "TODO",
			Action: func(context *cli.Context) error {
				// TODO
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
