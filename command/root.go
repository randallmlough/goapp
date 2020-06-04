package command

import (
	"fmt"
	"github.com/randallmlough/goapp/adapters/colors"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func Execute() error {
	app := cli.NewApp()
	app.Name = "goapp"
	app.Version = "0.0.1"

	app.Before = LoadConfig
	app.Commands = []*cli.Command{
		NewCmd,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return nil
}

func Println(msg string, task interface{}) {
	fmt.Printf("%s %s %v\n", colors.Yellow("[gogen]"), colors.Magenta(msg), task)
}
func PrintErr(msg string, err error) {
	fmt.Printf("%s %s %s\n", colors.Yellow("[gogen]"), colors.Black(msg).BgRed(), err)
}
