package main

import (
	"fmt"
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
func emptifier(list []string, val string) string {
	if val == "" {
		return list
	}
	return append(list, val)
}
func main() {
	wd := cwdir()
	fmt.Print(wd)
	emp := make([]string, "njan", "nee")
	fmt.Print(emptifier(emp, wd))
}
