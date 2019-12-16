package mods

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/talal/go-bits/color"
)

// getDir returns information regarding the current working directory.
func getDir(cwd string) string {
	if cwd == "/" {
		return color.Sprintf(color.Blue, cwd)
	}

	nearestAccessiblePath := findNearestAccessiblePath(cwd)

	if nearestAccessiblePath != cwd {
		inAccessiblePath := strings.TrimPrefix(cwd, nearestAccessiblePath)
		//nearestAccessiblePath = shortenLongPath(stripHomeDir(nearestAccessiblePath), 1)

		return color.Sprintf(color.Blue, nearestAccessiblePath) +
			color.Sprintf(color.Red, inAccessiblePath)
	}

	pathToDisplay := stripHomeDir(cwd)
	//pathToDisplay = shortenLongPath(pathToDisplay, 2)
	gbpath := pathToDisplay[strings.LastIndex(pathToDisplay, "/")+1:]
	gitDir, err := findGitRepo(cwd)
	handleError(err)
	isconmod := iscontentmodified(gitDir)

	if gitDir != "" {
		status := make(chan string)
		go findstatus(isconmod, gbpath, gitDir, status)
		imod := <-status
		return imod
	}
	return color.Sprintf(color.Cyan, pathToDisplay)
}

func findNearestAccessiblePath(path string) string {
	_, err := os.Stat(path)
	if err == nil {
		return path
	}

	return findNearestAccessiblePath(filepath.Dir(path))
}

func findstatus(mods ismodified, path string, gdir string, status chan string) {

	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.Magenta, path) + " on " +
		color.Sprintf(color.Green, ` `+branch)
	if mods.notStaged != 0 && mods.untracked != 0 {
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`!]`+`[`+strconv.Itoa(mods.untracked)+`+]`)
	}
	if mods.notStaged != 0 {
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`!]`)
	}
	if mods.untracked != 0 {
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.untracked)+`+]`)
	}
	cmod := make(chan int)

	go isahead(path, branch, cmod)
	ahead := <-cmod
	fmt.Print(ahead)
	if ahead != 0 {
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(ahead)+`x]`)
	}

	status <- nm
}
