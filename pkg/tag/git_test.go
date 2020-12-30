package git

import "testing"

func TestGenerateSHAForDirectory(t *testing.T) {
	// t.Err/orf("Failed to generate a SHA for a git directory")
	Tag(".")
}

func TestGenerateSHAForFile(t *testing.T) {
	// t.Errorf("Failed to generate a SHA for a git file")
}
