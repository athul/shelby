package mods

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

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

	refSpecDisplay := strings.TrimPrefix(refSpec, "ref: refs/heads/")
	return refSpecDisplay
}
