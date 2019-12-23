package mods

import (
	"strconv"

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
	status <- nm
}
