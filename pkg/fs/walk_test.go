package fs

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	found := true
	files, _ := Walk("test")
	for _, f := range files {
		fmt.Println(f)
		found = found && (f == "walk" ||
			f == "test" ||
			f == "walk/sub" ||
			f == "walk/sub/three.txt" ||
			f == "walk/one.txt")
	}

	if !found {
		t.Errorf("Couldn't locate expected files")
	}
}
