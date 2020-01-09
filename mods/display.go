package mods

import (
	"fmt"
	"strconv"

	"github.com/talal/go-bits/color"
)

/* func findstatus(mods ismodified, path string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightYellow,path) + " on " + color.Sprintf(color.BrightGreen, ` `+branch)
	switch {
	case mods.notStaged != 0 && mods.untracked != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`+]`+`[`+strconv.Itoa(mods.untracked)+`!]`)
	case mods.notStaged != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.notStaged)+`+]`)
	case mods.untracked != 0:
		status <- nm + color.Sprintf(color.BrightRed, ` [`+strconv.Itoa(mods.untracked)+`!]`)
	case mods.notStaged != 0 && mods.untracked != 0 && mods.
	default:
		status <- nm
	}
} */

func dispstats(m ismodified, path string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightYellow, path) + " on " + color.Sprintf(color.BrightGreen, ` `+branch)
	states := map[string]string{
		"ahead":  "↑",
		"behind": "↓",
		"both":   "⇅",
	}
	ius := "!"
	itr := "+"
	nstg_count := strconv.Itoa(m.notStaged)
	ntrc_count := strconv.Itoa(m.untracked)
	stt := states[m.state]

	if m.utrbool == true && m.ustbool == true {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s][%s%s] %s", nstg_count, ius, ntrc_count, itr, stt))
	}
	if m.utrbool == false && m.ustbool == true {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s] %s", nstg_count, ius, stt))
	}
	if m.utrbool == true && m.ustbool == false {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s] %s", ntrc_count, itr, stt))
	}
	if m.ustbool == false && m.utrbool == false {
		status <- nm + color.Sprintf(color.BrightRed, stt)
	}

}
