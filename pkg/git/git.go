package hack

import (
	"bytes"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// GitUtil represents git actions
type GitUtil struct {
	EntryPath    string
	TopLevelPath string
}

// NewGitUtil returns an instance of GitUtil
// with TopLevelPath set to the parent of the entryPath
// that contains the .git folder
func NewGitUtil(entryPath string) (*GitUtil, error) {
	var err error
	util := GitUtil{EntryPath: entryPath}
	util.TopLevelPath, err = util.ExecCommand("rev-parse", "--show-toplevel")
	return &util, err
}

// AppendDirty will check the status of git and append -dirty
// to the provided string if uncommited changes were found
func (g *GitUtil) AppendDirty(sha string) (string, error) {
	status, err := g.ExecCommand("status", ".", "--porcelain")

	if err != nil {
		return sha, err
	}

	if len(status) != 0 {
		return fmt.Sprintf("%v-dirty", sha), nil
	}

	return sha, nil
}

// GetCommitSha returns the git commit from the command
// git rev-list -1 HEAD
func (g *GitUtil) GetCommitSha() (string, error) {
	return g.ExecCommand("rev-list", "-1", "HEAD")
}

// GetShortCommitSha returns the git commit from the command
// git rev-list -1 HEAD --abbrev-commit
func (g *GitUtil) GetShortCommitSha() (string, error) {
	return g.ExecCommand("rev-list", "-1", "HEAD", "--abbrev-commit")
}

// ExecCommand returns the output of a git command
func (g *GitUtil) ExecCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = g.TopLevelPath

	// Capture output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run
	err := cmd.Run()

	if err != nil {
		log.Debugln(stderr.String())
	}

	return string(bytes.TrimSpace(out.Bytes())), err
}
