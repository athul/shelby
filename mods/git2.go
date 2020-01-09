package mods

/* package mods

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	dir, err := os.Getwd()
	r, err := git.PlainOpen(dir)
	rem, err := r.Remote("origin")
	r1 := rem.String()
	w, err := r.Worktree()
	s, err := w.Status()
	es := s.String()
	if strings.ContainsAny(es, "?") {
		fmt.Print("Yes ?? is there")
	}

	if err != nil {
		panic(err)
	}
	fmt.Print("\n" + r1)
}

type repstats struct {
	ahead      int
	behind     int
	untracked  int
	notStaged  int
	staged     int
	conflicted int
	stashed    int
	remote     string
}

func getstats(dir string) repstats {
	statevars := repstats{}
	dir, err := os.Getwd()
	repo, err := git.PlainOpen(dir)
	rem, err := repo.Remote("origin")
	wtree, err := repo.Worktree()
	stats, err := wtree.Status()
	statsrings := stats.String()
	r1 := rem.String()
	handleError(err)


}
*/
