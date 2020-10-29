package docker

import (
	"os/exec"
	"testing"
)

func TestImageExists(t *testing.T) {
	const testImage = "alpine:latest"

	// Grab an image from the cli
	cmd := exec.Command("docker", "pull", testImage)
	if err := cmd.Run(); err != nil {
		t.Errorf("Couldn't pull %v from the cli", testImage)
	}

	// Attempt to grab an image
	exists, err := ImageExists(testImage)

	if err != nil {
		t.Error(err)
	}

	if !exists {
		t.Errorf("Couldn't pull %v in app", testImage)
	}
}
