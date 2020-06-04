package command

import (
	"github.com/urfave/cli/v2"
)

var NewCmd = &cli.Command{
	Name:    "new",
	Aliases: []string{"n"},
	Usage:   "creates something new",
	Subcommands: []*cli.Command{
		NewProjectCmd,
	},
	//Flags: []cli.Flag{
	//	&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
	//	&cli.StringFlag{Name: "config, c", Usage: "the config filename"},
	//},
}
