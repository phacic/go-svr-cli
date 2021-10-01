package main

import (
	cli "github.com/urfave/cli/v2"
)

// ServerCommand is the cli command that run the server
func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{},
		Usage:   "Start the server",
		Action: func(c *cli.Context) error {
			RunServer()
			return nil
		},
	}
}

// UpCommand is the cli command for counting up
func UpCommand() *cli.Command {
	return &cli.Command{
		Name:    "up",
		Aliases: []string{"u"},
		Usage:   "Count up to stop value.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "stop",
				Aliases: []string{"s"},
				Value:   10,
			},
		},
		Action: func(c *cli.Context) error {
			stop := c.Int("stop")
			_, err := MakeRequest(stop, stop)
			return err

		},
	}
}

// DownCommand is the cli command for counting down
// DownCommand is the cli command for counting down
func DownCommand() *cli.Command {
	return &cli.Command{
		Name:    "down",
		Aliases: []string{"d"},
		Usage:   "Count down from value",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "start",
				Aliases: []string{"s"},
				Value:   10,
			},
		},
		Action: func(c *cli.Context) error {
			stop := c.Int("start")
			_, err := MakeRequest(stop, stop)
			return err

		},
	}
}
