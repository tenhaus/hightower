package docker

import (
	"context"
	"testing"
)

func TestBuild(t *testing.T) {
	const testImage = "buildtest:test"
	ctx := context.Background()

	Build(&ctx, BuildOptions{
		Path:       "testinfo/base",
		Dockerfile: "testinfo/base/Dockerfile",
		Tag:        testImage,
		Cache:      false,
	})

	exists, err := ImageExists(testImage)

	if err != nil || !exists {
		t.Errorf("Failed buiding %v", testImage)
	}

}
