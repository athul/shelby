package mods

import (
	"os"
	"path/filepath"
	"strconv"
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
	sshtrue := isssh()
	name, host := gethost()
	if sshtrue == true {
		return strings.Replace(path, os.Getenv("HOME"), color.Sprintf(color.BrightGreen, name+" on "+host)+" ~", 1)
	}
	return strings.Replace(path, os.Getenv("HOME"), "~", 1)
}
func handleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.BrightRed, "Prompt error: %v\n", err)
	}
}

// Return the Exit Code if an arg is passed
func getExitCodeField(arg string) string {
	exitCode, err := strconv.Atoi(arg)
	if err == nil && exitCode > 0 {
		return color.Sprintf(color.BrightRed, "exit:"+arg)
	}
	return ""
}
