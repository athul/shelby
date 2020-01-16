package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/athul/shelby/mods"
	"github.com/urfave/cli"
)

var shell = path.Base(os.Getenv("SHELBY_SHELL"))

/* func main() {
	fmt.Printf(shell)

	if len(os.Args) > 1 {
		color.Fprint(os.Stderr, color.Red, "Error.....\n")
		os.Exit(1)
	}
	os.Stdout.Write([]byte("\n" + mods.Info() + "\n"))
}.
*/
func main() {
	app := cli.NewApp()
	app.Usage = "The Prompt with Wings"
	app.Commands = []cli.Command{
		sbInit,
	}
	app.Run(os.Args)
}

var sbInit = cli.Command{
	Name:        "init",
	Usage:       "Initialize the Shell",
	Description: `Put 'eval "$(shelby init)"' in your ~/.profile or ~/.bashrc or ~/.zshrc`,
	/* Action: func(ctx *cli.Context) {
		script := mods.UseShell(shell)
		//fmt.Print(script)
		os.Stdout.Write([]byte("\n" + mods.Info() + "\n"))
		os.Stdout.Write([]byte(script))
	}, */

	Action: func(ctx *cli.Context) error {
		script, err := ioutil.ReadFile("/usr/local/bin/init." + shell)
		if err == nil {
			//fmt.Print(script)
			os.Stdout.Write([]byte("\n" + mods.Info() + "\n"))
			os.Stdout.Write(script)
		} else {
			mods.HandleError(err)
		}
		return nil
	},
}
