package mods

import (
	"errors"
	"strings"
)

var cwd string

// Info returns the prompt info line.
func Info() string {
	cwd = cwdir()
	if cwd == "" {
		handleError(errors.New("could not get path for current working directory"))
	}

	info := make([]string, 1)
	info = emptifier(info, getDir(cwd))
	return strings.Join(info, " ")

}
