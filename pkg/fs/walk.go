package fs

import (
	"log"
	"os"
	"path/filepath"
)

// Walk returns a list of all files in the given directory
func Walk(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, path+info.Name())

		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return files, err
}
