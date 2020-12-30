package hack

import (
	"bytes"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// GitTopLevel finds the nearest parent with a .git folder
func GitTopLevel(path string) (string, error) {

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = path

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
