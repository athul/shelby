package mods

import (
	"os"
	"os/user"
	"path"
	"strings"
)

// All these functins are used in home_functions.go
// in the stripHomeDir() fn

// Get the Hostname for SSH
func gethost() (string, string) {
	hname, _ := os.Hostname()
	uname := os.Getenv("USER")
	return uname, strings.SplitN(hname, ".", 2)[0]
}

// Check if the system is on SSH
// using the SSH_CLIENT environment variable
func isssh() bool {
	cmd := os.Getenv("SSH_CLIENT")
	if cmd != "" {
		return true
	}
	return false
}

// Return if any Virtual ENvironment is on by
// checking the VIRTUAL_ENV variable is present.
// Only shows the venv on a Git Repository
func getifvenv() string {
	env := os.Getenv("VIRTUAL_ENV")
	return path.Base(env)
}

// Get Go Path is the the GOPATH environment variable is same
// as current working directory
func getgopath() bool {
	gpath := os.Getenv("GOPATH")
	if strings.Contains(os.Getenv("PWD"), gpath) {
		return true
	}
	return false
}

// Checks if the User is on root user.
// `sudo su` and we'll be logged into root mode
// For root the UID will be equal to 0.
func checkifRoot() bool {
	currentUser, _ := user.Current()
	if currentUser.Uid == "0" {
		return true
	}
	return false
}
