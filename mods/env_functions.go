package mods

import (
	"os"
	"path"
	"strings"
)

// Get the Hostname for SSH
func gethost() (string, string) {
	hname, _ := os.Hostname()
	uname := os.Getenv("USER")
	return uname, strings.SplitN(hname, ".", 2)[0]
}

// Check if the system is on SSH
func isssh() bool {
	cmd := os.Getenv("SSH_CLIENT")
	return cmd != ""
}

// Return if any Virtual ENvironment is on
func getifvenv() string {
	env := os.Getenv("VIRTUAL_ENV")
	return path.Base(env)
}
func getgopath() bool {
	gpath := os.Getenv("GOPATH")
	return strings.Contains(os.Getenv("PWD"), gpath)
}
