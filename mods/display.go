package mods

import (
	"os"
	"strconv"
	"strings"

	"github.com/talal/go-bits/color"
)

func findstatus(mods ismodified, path string, gdir string, status chan string) {
	branchchannel := make(chan string)
	go currentGitBranch(gdir, branchchannel)
	branch := <-branchchannel
	nm := color.Sprintf(color.BrightYellow, path) + " on " + color.Sprintf(color.BrightGreen, ` `+branch)
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
	//fmt.Print(ahead)
	if ahead != 0 {
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(ahead)+`x]`)
	}

	status <- nm
}
func gethost() string {
	hname, _ := os.Hostname()
	return strings.SplitN(hname, ".", 2)[0]
}
func isssh(probs chan bool) {
	cmd := os.Getenv("SSH_CLIENT")
	if cmd != "" {
		probs <- true
	}
	probs <- false
}
