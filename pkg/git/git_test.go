package hack

import (
	"os"
	"strings"
	"testing"
)

func initGit(t *testing.T) *GitUtil {
	path, err := os.Getwd()

	if err != nil {
		t.Error("Could not get the current working directory")
	}

	git, err := NewGitUtil(path)

	if err != nil {
		t.Error(err)
	}

	return git
}

func TestGitTopLevel(t *testing.T) {
	git := initGit(t)

	if !strings.Contains(git.EntryPath, git.TopLevelPath) {
		t.Errorf("Expected %v to contain %v", git.EntryPath, git.TopLevelPath)
	}

	if len(git.TopLevelPath) == 0 {
		t.Errorf("Top level path is nil")
	}
}

func TestGetCommitSha(t *testing.T) {
	git := initGit(t)
	sha, err := git.GetCommitSha()

	if err != nil {
		t.Error(err)
	}

	if len(sha) == 0 {
		t.Errorf("Commit sha is nil")
	}

	if len(sha) <= 7 {
		t.Errorf("Commit sha length is incorrect")
	}
}

func TestGetShortCommitSha(t *testing.T) {
	git := initGit(t)
	sha, err := git.GetShortCommitSha()

	if err != nil {
		t.Error(err)
	}

	if len(sha) == 0 {
		t.Errorf("Commit sha is nil")
	}

	if len(sha) != 7 {
		t.Errorf("Commit sha length is incorrect")
	}
}

func TestAppendDirty(t *testing.T) {
	git := initGit(t)
	sha, err := git.GetShortCommitSha()

	if err != nil {
		t.Error(err)
	}

	// Move a file so git status displays
	// changes
	os.Rename("testdata/dirty", "testdata/dirtytest")
	defer os.Rename("testdata/dirtytest", "testdata/dirty")

	dirtySha, err := git.AppendDirty(sha)

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(dirtySha, "-dirty") {
		t.Errorf("Failed to append -dirty")
	}
}
