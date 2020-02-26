package mods

import "github.com/talal/go-bits/color"

//Display the status of the Git Repo
func dispstats(m ismodified, path string, gdir string, status chan string) {
	branch := currentGitBranch(gdir)
	nm := color.Sprintf(color.BrightGreen, path) + " on " + color.Sprintf(color.BrightYellow, ` `+branch)
	states := map[string]string{
		"ahead":    "↑",
		"behind":   "↓",
		"diverged": "⇅",
	}
	ius := "!"
	itr := "+"
	stg := "✔"
	//nstgcount := strconv.Itoa(m.notStaged)
	//ntrccount := strconv.Itoa(m.untracked)
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
