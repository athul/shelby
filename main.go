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
