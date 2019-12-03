package mods

import (
	"errors"
	"os"
	"strings"
)

// Info returns the prompt info line.
func Info(args []string) int {
	cwd := cwdir()
	if cwd == "" {
		handleError(errors.New("could not get path for current working directory"))
	}

	info := make([]string, 1)
	info = emptifier(info, getDir(cwd))
	if len(args) > 0 {
		info = emptifier(info, getExitCodeField(args[0]))
	}
	line := strings.Join(info, " ")
	os.Stdout.Write([]byte(line + "\n"))

	//print second line: a letter identifying the shell, and the final "$ ")
	shellIdent := ""
	os.Stdout.Write([]byte(shellIdent + " "))

	return 0

	//return strings.Join(info, " ")

}
