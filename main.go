package main

import (
	"os"

	"github.com/athul/shelby/mods"
	"github.com/talal/go-bits/color"
)

func main() {
	//os.Exit(shelbyexec(filepath.Base(os.Args[0]), os.Args[1:], true))

	if len(os.Args) > 1 {
		color.Fprint(os.Stderr, color.Red, "Error.....\n")
		os.Exit(1)
	}
	os.Stdout.Write([]byte("\n" + mods.Info() + "\n"))
	shellIdent := ""
	os.Stdout.Write([]byte(shellIdent + "❯ "))
}

/* func shelbyexec(applet string, args []string, allowGofu bool) int {
	//allow explicit specification of applet as "./build/gofu <applet> <args>"
	if allowGofu && applet == "main" || allowGofu && applet == "shelby" {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Working I guess")
			return 1
		}
		return shelbyexec(args[0], args[1:], false)
	}

	switch applet {
	case "run":
		return mods.Info(args)
	default:
		fmt.Fprintln(os.Stderr, "ERROR: unknown applet: "+applet)
		return 255
	}
} */
