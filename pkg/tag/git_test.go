package git

import (
	"log"
	"os"
	"testing"
)

func TestGenerateSHAForDirectory(t *testing.T) {

	path, err := os.Getwd()

	if err != nil {
		t.Error("Could not get the current working directory")
	}

	tag, err := Tag(path)

	if err != nil {
		t.Error(err)
	}

	log.Println(tag)
}
