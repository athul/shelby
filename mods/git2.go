package mods

import (
	"fmt"
	"os"

	git "gopkg.in/libgit2/git2go.v27"
)

func main() {
	dir, err := os.Getwd()
	r, err := git.OpenRepository(dir)
	/* 	ref, err1 := r.Head()
	   	branch := plumbing.Str
		   rem, err2 := r.Remotes() */
	rem, err2 := r.Remotes.Lookup("origin")
	url := rem.Url()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(r, url, rem, err2)
}

//func getbranch(path string)
