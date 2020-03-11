package main

import (
	"os"

	"github.com/athul/shelby/mods"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "The Prompt with Wings"
	app.Commands = []cli.Command{
		sbInit,
		sbInfo,
	}
	app.Run(os.Args)
}

var sbInit = cli.Command{
	Name:        "init",
	Usage:       "Initialize the Shell",
	Description: `Put 'eval "$(shelby init <shell>)"' in your ~/.profile or ~/.bashrc or ~/.zshrc`,
	Action: func(ctx *cli.Context) error {
		if ctx.NArg() == 1 {
			script := mods.UseShell(ctx.Args().First())
			if script != "" {
				os.Stdout.Write([]byte(script))
			} else {
				os.Stderr.Write([]byte("Error.....\n"))
				os.Exit(1)
			}
		}
		return nil
	},
}

var sbInfo = cli.Command{
	Name:        "info",
	Usage:       "Print directory info",
	Description: "Print directory and other mods info in a line",
	Action: func(ctx *cli.Context) error {
		os.Stdout.Write([]byte(mods.Info()))
		return nil
	},
}
