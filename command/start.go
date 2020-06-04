package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

var (
	StartCmd = &cli.Command{
		Name:   "start",
		Action: startHandler,
		Before: altsrc.InitInputSourceWithContext(startFlags, altsrc.NewYamlSourceFromFlagFunc("config")),
		Flags:  startFlags,
	}
	startFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "path to the configuration file",
		},
		altsrc.NewStringFlag(
			&cli.StringFlag{
				Name:  "listen-addr",
				Value: "127.0.0.9",
				Usage: "listening address",
			},
		),
	}
)

func startHandler(c *cli.Context) error {
	fmt.Println("CONTEXT", configFromContext(c))
	fmt.Println("listen-addr: ", c.String("listen-addr"))
	return nil
}
