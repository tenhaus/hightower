package fs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Walk returns a list of all files in the given directory
func Walk(dir string, includeDir bool) ([]string, error) {
	var files []string
	cleanDir := dir

	// Strip out ./ from the beginning of path
	// if it exists
	if strings.Index(dir, "./") == 0 {
		cleanDir = dir[2:]
	}

	// Build a list of files from cleanDir
	err := filepath.Walk(cleanDir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		// We don't have to include the root directory
		if cleanDir == path {
			return nil
		}

		// Don't include directories
		if !includeDir && info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files, err
}
