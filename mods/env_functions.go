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
	if cmd != "" {
		return true
	}
	return false
}

// Return if any Virtual ENvironment is on
func getenv() string {
	env := os.Getenv("VIRTUAL_ENV")
	return path.Base(env)
}
