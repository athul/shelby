package mods

import (
	"strconv"

	"github.com/talal/go-bits/color"
)

func findstatus(mods ismodified, path string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightYellow, path) + " on " + color.Sprintf(color.BrightGreen, ` `+branch)
	switch {
	case mods.notStaged != 0 && mods.untracked != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`+]`+`[`+strconv.Itoa(mods.untracked)+`!]`)
	case mods.notStaged != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`+]`)
	case mods.untracked != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.untracked)+`!]`)
	default:
		status <- nm
	}
}
