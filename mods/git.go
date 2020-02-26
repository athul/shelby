package mods

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var dir string = cwdir()

type ismodified struct {
	utrbool   bool
	ustbool   bool
	untracked int
	notStaged int
	staged    bool
	state     string
}

// Find a git repo if a .git folder is found
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

// Find the Current Git Branch using git files
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

// Setting the envs for Git Command Execution
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

// Run git Commands
func rungitcommands(cmd string, args ...string) (string, error) { //  as params
	command := exec.Command(cmd, args...) //,  as params
	command.Env = gitProcessEnv()
	op, err := command.Output()
	return string(op), err
}

// Parse the Git Status for untracked, modified and staged
func parseGitStats(status []string) ismodified {
	stats := ismodified{}
	if len(status) > 1 {
		for _, line := range status[1:] {
			if len(line) > 2 {
				code := line[:2]
				switch code {
				case "??":
					stats.untracked++
					stats.utrbool = true
				default:
					if code[0] != ' ' {
						stats.staged = true
					}
					if code[1] != ' ' {
						stats.notStaged++
						stats.ustbool = true
					}
				}
			}
		}
	}
	return stats
}

// Check if the content is Modified for a git repo
func iscontentmodified(path string) ismodified {
	out, err := rungitcommands("git", "status", "--porcelain", "-b", "--ignore-submodules")
	status := strings.Split(out, "\n")
	stats := parseGitStats(status)
	if strings.Contains(out, ",") {
		stats.state = "diverged"
	} else if strings.Contains(out, "behind") {
		stats.state = "behind"
	} else if strings.Contains(out, "ahead") {
		stats.state = "ahead"
	} else {
		stats.state = "clean"
	}
	if err != nil {

	}
	return stats

}
