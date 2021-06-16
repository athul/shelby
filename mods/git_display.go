package mods

import (
	"path"

	"github.com/talal/go-bits/color"
)

//Display the status of the Git Repo
func dispstats(m ismodified, fdpath string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightGreen, fdpath) + " on " + color.Sprintf(color.BrightYellow, ` `+branch)
	if getgopath() {
		nm = color.Sprintf(color.BrightMagenta, path.Base(fdpath)) + " on " + color.Sprintf(color.BrightYellow, ` `+branch)
	}
	ius := "!" //icon for unstaged
	itr := "+" //icon for tracked
	stg := "✔" //icon for staged
	states := map[string]string{
		"ahead":    "↑",
		"behind":   "↓",
		"diverged": "⇅",
	}
	stt := states[m.state]
	if m.utrbool && m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, " [%d%s][%d%s] %s", m.notStaged, ius, m.untracked, itr, stt)
	}
	if !m.utrbool && m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, " [%d%s] %s", m.notStaged, ius, stt)
	}
	if m.utrbool && !m.ustbool {
		status <- nm + color.Sprintf(color.BrightRed, " [%d%s] %s", m.untracked, itr, stt)
	}
	if !m.ustbool && !m.utrbool && m.staged {
		status <- nm + color.Sprintf(color.BrightMagenta, " %s %s ", stg, stt)
	}
	if !m.ustbool && !m.utrbool {
		status <- nm + color.Sprintf(color.BrightBlue, " %s", stt)
	}
}
