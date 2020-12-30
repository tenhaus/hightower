package git

import (
	"log"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

// Tag ...
func Tag(path string) {

	// Initialize a git object
	fs := osfs.New(path)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		if err != nil {
			log.Fatal("Could not resolve git")
		}
	}

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
	_, err := git.Open(s, fs)

	if err != nil {
		log.Fatal("Could not open git repo")
	}

	defer s.Close()

}
