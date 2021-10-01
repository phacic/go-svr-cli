package main

import (
	"os"

	cli "github.com/urfave/cli/v2"
)

func RunCli() {
	var verbositiy bool

	app := cli.NewApp()
	app.Usage = "Counter cli"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "verbose",
			Aliases:     []string{"v"},
			Usage:       "",
			Destination: &verbositiy,
		},
	}

	app.Commands = []*cli.Command{
		ServerCommand(),
		UpCommand(),
		DownCommand(),
	}

	err := app.Run(os.Args)
	FatalOnErr(err)

}

func main() {
	RunCli()
}
