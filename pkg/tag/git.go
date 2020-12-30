package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/tenhaus/hightower/pkg/hack"
)

// Tag ...
func Tag(path string) (string, error) {

	// Find our top level .git folder
	topLevelPath, err := hack.GitTopLevel(path)
	if err != nil {
		return "", err
	}

	topLevelPath = fmt.Sprintf("%v/.git", topLevelPath)

	// Initialize a git object
	fs := osfs.New(topLevelPath)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		if err != nil {
			log.Fatal("Could not resolve git")
		}
	}

	// Open the repo
	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
	r, err := git.Open(s, fs)

	if err != nil {
		log.Fatal("Could not open git repo")
	}

	defer s.Close()

	// Get head
	ref, err := r.Head()

	if err != nil {
		return "", err
	}

	w, err := r.Worktree()
	status, err := w.Status()
	fmt.Println(ref.Hash())
	if status.IsClean() {
		fmt.Println("clean")
	} else {
		fmt.Println("dirty")
	}

	return "", nil

}
