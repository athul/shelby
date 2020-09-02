package mods

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

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

// Replace the Path in Prompt with different formatting
// with different functions
// isssh() checks if you're in SSH
// checkifRoot() checks if you're the root user
// getgopath() check if you're in the Go path
func stripHomeDir(path string) string {
	name, host := gethost()
	if isssh() {
		return strings.Replace(path, os.Getenv("HOME"), color.Sprintf(color.BrightGreen, name+" on "+host)+" ~", 1)
	}
	if checkifRoot() {
		return strings.Replace(path, "", color.Sprintf(color.BrightRed, "root")+" ~", 1)
	}
	if getgopath() {
		return strings.Replace(color.Sprintf(color.BrightMagenta, path), os.Getenv("HOME"), "🐭 ~", 1)
	}

	return strings.Replace(path, os.Getenv("HOME"), "~", 1)
}

//HandleError will handle all errors
func HandleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.BrightRed, "Prompt error: %v\n", err)
	}
}
