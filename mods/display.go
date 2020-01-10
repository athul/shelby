package mods

import (
	"fmt"
	"strconv"

	"github.com/talal/go-bits/color"
)

func dispstats(m ismodified, path string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightGreen, path) + " on " + color.Sprintf(color.BrightYellow, ` `+branch)
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

	if m.utrbool && m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s][%s%s] %s", nstg_count, ius, ntrc_count, itr, stt))
	}
	if !m.utrbool && m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s] %s", nstg_count, ius, stt))
	}
	if m.utrbool && !m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, fmt.Sprintf(" [%s%s] %s", ntrc_count, itr, stt))
	}
	if !m.ustbool && !m.utrbool {
		status <- nm + color.Sprintf(color.BrightBlue, " "+stt)
	}

}
