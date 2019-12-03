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

	info := make([]string, 0, 3)
	info = emptifier(info, getDir(cwd))
	return strings.Join(info, " ")
}
