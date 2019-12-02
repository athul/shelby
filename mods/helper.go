package mods

import (
	"os"
	"path/filepath"
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

//Test Cases for debugging are below
/* func main() {
	wd := cwdir()
	fmt.Print(wd)
	emp := make([]string, 0, 3)
	fmt.Println(emptifier(emp, wd))
} */
