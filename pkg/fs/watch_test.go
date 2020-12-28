package fs

import "testing"

func TestWatchFile(t *testing.T) {
	t.Errorf("Failed to detect changes to a single file")
}

func TestWatchDirectory(t *testing.T) {
	t.Errorf("Failed to detect changes to a directory")
}
