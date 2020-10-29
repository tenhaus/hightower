package fs

import (
	"testing"
)

func TestWalkFiles(t *testing.T) {
	found := true
	files, _ := Walk("./testdata", false)
	for _, f := range files {
		found = found && (f.Path == "testdata/walk/sub/subsub/two.txt" ||
			f.Path == "testdata/walk/sub/three.txt" ||
			f.Path == "testdata/walk/one.txt")
	}

	if !found {
		t.Errorf("Couldn't find expected files")
	}
}

func TestWalkFilesAndDirectories(t *testing.T) {
	found := true
	files, _ := Walk("./testdata", true)
	for _, f := range files {
		found = found && (f.Path == "testdata" ||
			f.Path == "testdata/walk" ||
			f.Path == "testdata/walk/sub" ||
			f.Path == "testdata/walk/sub/subsub" ||
			f.Path == "testdata/walk/sub/subsub/two.txt" ||
			f.Path == "testdata/walk/sub/three.txt" ||
			f.Path == "testdata/walk/one.txt")
	}

	if !found {
		t.Errorf("Couldn't find expected files")
	}
}
