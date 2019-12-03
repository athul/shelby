package mods

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/talal/go-bits/color"
)

func cwdir() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = os.Getenv("PWD")
	}
	return filepath.Clean(wd)
}
func emptifier(list []string, val string) []string {
	if val == "" {
		return list
	}
	return append(list, val)
}
func stripHomeDir(path string) string {
	return strings.Replace(path, os.Getenv("HOME"), "~", 1)
}
func handleError(err error) {
	if err != nil {
		color.Fprintf(os.Stderr, color.Red, "Prompt error: %v\n", err)
	}
}

//Test Cases for debugging are below
/* func main() {
	wd := cwdir()
	fmt.Print(wd)
	emp := make([]string, 0, 3)
	fmt.Println(emptifier(emp, wd))
} */
