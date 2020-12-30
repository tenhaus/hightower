package hack

import (
	"os"
	"strings"
	"testing"
)

func TestGitTopLevel(t *testing.T) {
	path, err := os.Getwd()

	if err != nil {
		t.Error("Could not get the current working directory")
	}

	topLevelPath, err := GitTopLevel(path)

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(path, topLevelPath) {
		t.Errorf("Expected %v to contain %v", path, topLevelPath)
	}
}
