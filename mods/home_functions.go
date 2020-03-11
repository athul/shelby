package mods

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

//var host = gethost()
// Get the current working Directory
func cwdir() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = os.Getenv("PWD")
	}
	return filepath.Clean(wd)
}

// Clean the array
func emptifier(list []string, val string) []string {
	if val == "" {
		return list
	}
	return append(list, val)
}

//
func stripHomeDir(path string) string {
	name, host := gethost()
	if isssh() {
		return strings.Replace(path, os.Getenv("HOME"), color.Sprintf(color.BrightGreen, name+" on "+host)+" ~", 1)
	}
	return strings.Replace(path, os.Getenv("HOME"), "~", 1)
}

//HandleError will handle all errors
func HandleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.BrightRed, "Prompt error: %v\n", err)
	}
}
