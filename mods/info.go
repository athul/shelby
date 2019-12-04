package mods

import (
	"errors"
	"strings"
)

// Info returns the prompt info line.
func Info() string {
	cwd := cwdir()
	if cwd == "" {
		handleError(errors.New("could not get path for current working directory"))
	}

	info := make([]string, 1)
	info = emptifier(info, getDir(cwd))
	//line := strings.Join(info, " ")
	//os.Stdout.Write([]byte(line + "\n"))

	//print second line: a letter identifying the shell, and the final "$ ")

	//return 0

	return strings.Join(info, " ")

}
