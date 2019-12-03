package mods

import (
	"os"
	"path/filepath"
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
		nearestAccessiblePath = shortenLongPath(stripHomeDir(nearestAccessiblePath), 1)

		return color.Sprintf(color.Blue, nearestAccessiblePath) +
			color.Sprintf(color.Red, inAccessiblePath)
	}

	pathToDisplay := stripHomeDir(cwd)
	pathToDisplay = shortenLongPath(pathToDisplay, 2)

	gitDir, err := findGitRepo(cwd)
	handleError(err)

	if gitDir != "" {
		return color.Sprintf(color.Red, pathToDisplay) + " " +
			color.Sprintf(color.Green, ``+currentGitBranch(gitDir))
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
func shortenLongPath(path string, length int) string {
	pList := strings.Split(path, "/")
	if len(pList) < 7 {
		return path
	}
	lname := pList[len(pList)-1]
	shortenedPList := pList[:len(pList)-length]
	for i, v := range shortenedPList {
		// shortenedPList[0] will be an empty string due to leading '/'
		if len(v) > 0 {
			shortenedPList[i] = v[:1]
		}
	}

	shortenedPList = append(shortenedPList, pList[len(pList)-length:]...)
	return lname
}
