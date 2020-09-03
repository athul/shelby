package mods

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

/*
This file does all the printing. Directories, git branches and whatnot
*/

// getDir returns information regarding the current working directory.
func getDir(cwd string) string {
	if cwd == "/" {
		return color.Sprintf(color.BrightCyan, cwd)
	}

	nearestAccessiblePath := findNearestAccessiblePath(cwd)

	if nearestAccessiblePath != cwd {
		inAccessiblePath := strings.TrimPrefix(cwd, nearestAccessiblePath)

		return color.Sprintf(color.Blue, nearestAccessiblePath) +
			color.Sprintf(color.Red, inAccessiblePath)
	}

	pathToDisplay := stripHomeDir(cwd)
	gitDir, err := findGitRepo(cwd)
	HandleError(err)
	env := getifvenv()
	if checkifRoot() {
		return color.Sprintf(color.BrightRed, "root ") + "at " + color.Sprintf(color.BrightCyan, pathToDisplay)
	}
	if gitDir != "" && env != "" && env != "." {
		isconmod := iscontentmodified(gitDir)
		status := make(chan string)
		go dispstats(isconmod, pathToDisplay, gitDir, status)
		imod := <-status
		return imod + color.Sprintf(color.BrightBlack, "(%s)", env)
	}
	if gitDir != "" {
		isconmod := iscontentmodified(gitDir)
		status := make(chan string)
		go dispstats(isconmod, pathToDisplay, gitDir, status)
		imod := <-status
		return imod
	}

	return color.Sprintf(color.BrightCyan, pathToDisplay)
}

// findNearestAccessiblePath takes the last string after the splitting
// the path from the the path from os.Stat fn recursively
func findNearestAccessiblePath(path string) string {
	_, err := os.Stat(path)
	if err == nil {
		return path
	}

	return findNearestAccessiblePath(filepath.Dir(path))
}
