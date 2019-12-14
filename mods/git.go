package mods

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type ismodified struct {
	ahead      int
	behind     int
	untracked  int
	notStaged  int
	staged     int
	conflicted int
	stashed    int
}

func findGitRepo(path string) (string, error) {
	gitEntry := filepath.Join(path, ".git")
	fi, err := os.Stat(gitEntry)
	switch {
	case err == nil:
		//found - continue below with further checks
	case !os.IsNotExist(err):
		return "", err
	case path == "/":
		return "", nil
	default:
		return findGitRepo(filepath.Dir(path))
	}

	if !fi.IsDir() {
		return "", nil
	}

	return gitEntry, nil
}
func currentGitBranch(gitDir string) string {
	bytes, err := ioutil.ReadFile(filepath.Join(gitDir, "HEAD"))
	if err != nil {
		handleError(err)
		return "unknown"
	}
	refSpec := strings.TrimSpace(string(bytes))

	// detached HEAD?
	if !strings.HasPrefix(refSpec, "ref: refs/") {
		return "detached"
	}

	branch := strings.TrimPrefix(refSpec, "ref: refs/heads/")
	return branch
}
func gitProcessEnv() []string {
	home, _ := os.LookupEnv("HOME")
	path, _ := os.LookupEnv("PATH")
	env := map[string]string{
		"LANG": "C",
		"HOME": home,
		"PATH": path,
	}
	result := make([]string, 0)
	for key, value := range env {
		result = append(result, fmt.Sprintf("%s=%s", key, value))
	}
	return result
}
func rungitcommands(cmd string, args ...string) (string, error) { //  as params
	command := exec.Command(cmd, args...) //,  as params
	command.Env = gitProcessEnv()
	op, err := command.Output()
	return string(op), err
}
func parseGitStats(status []string) ismodified {
	stats := ismodified{}
	if len(status) > 1 {
		for _, line := range status[1:] {
			if len(line) > 2 {
				code := line[:2]
				switch code {
				case "??":
					stats.untracked++
				case "DD", "AU", "UD", "UA", "DU", "AA", "UU":
					stats.conflicted++
				default:
					if code[0] != ' ' {
						stats.staged++
					}

					if code[1] != ' ' {
						stats.notStaged++
					}
				}
			}
		}
	}
	return stats
}
func iscontentmodified(path string) ismodified {
	out, err := rungitcommands("git", "status", "--porcelain", "-b", "--ignore-submodules")
	status := strings.Split(out, "\n")
	if err != nil {

	}
	stats := parseGitStats(status)
	brinfo:=
	if brinfo["local"] != "" {
		ahead, _ := strconv.ParseInt(brinfo["ahead"], 10, 32)
		stats.ahead = int(ahead)

		behind, _ := strconv.ParseInt(brinfo["behind"], 10, 32)
		stats.behind = int(behind)

		branch = brinfo["local"]
	} else {
		branch = getGitDetachedBranch(p)
	}
	return stats
}
