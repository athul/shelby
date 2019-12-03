package mods

import (
	"fmt"
	/* "strings"
	"ioutils" */
	"os"
	"path/filepath"
	//"github.com/talal/go-bits/colors"
)

type gitRepo struct {
	Rootpath string
	gDir     string
}

func getrepo(path string) (*gitRepo, error) {
	gitent := filepath.Join(path, ".git")
	fi, err := os.Stat(gitent)
	switch {
	case err == nil:
		//found - continue below with further checks
	case !os.IsNotExist(err):
		return nil, err
	case path == "/":
		return nil, nil
	default:
		return getrepo(filepath.Dir(path))
	}
	if fi.Mode().IsDir() {
		//normal case - .git is a directory
		return &gitRepo{Rootpath: path, gDir: gitent}, nil
	}
	return nil, fmt.Errorf("read %s: missing gitdir directive", gitent)
}
func main() {
	dir, ee := getrepo(filepath.Clean(os.Getenv("PWD")))
	fmt.Println(dir, ee)
}
