package fs

import (
	"testing"
)

func TestWalkFiles(t *testing.T) {
	found := true
	files, _ := Walk("./testdata", false)
	for _, f := range files {
		found = found && (f == "testdata/walk/sub/subsub/two.txt" ||
			f == "testdata/walk/sub/three.txt" ||
			f == "testdata/walk/one.txt")
	}

	if !found {
		t.Errorf("Couldn't locate expected files")
	}
}

func TestWalkFilesAndDirectories(t *testing.T) {
	found := true
	files, _ := Walk("./testdata", true)
	for _, f := range files {
		found = found && (f == "testdata" ||
			f == "testdata/walk" ||
			f == "testdata/walk/sub" ||
			f == "testdata/walk/sub/subsub" ||
			f == "testdata/walk/sub/subsub/two.txt" ||
			f == "testdata/walk/sub/three.txt" ||
			f == "testdata/walk/one.txt")
	}

	if !found {
		t.Errorf("Couldn't locate expected files")
	}
}
