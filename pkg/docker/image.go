package docker

import (
	"context"

	"github.com/docker/docker/api/types"
)

// ImageExists verifies that an image exists with the specific tag
func ImageExists(tag string) (bool, error) {

	ctx := context.Background()
	cli, err := NewClient()

	if err != nil {
		return false, err
	}

	// Get the list of images
	list, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		return false, err
	}

	// Check the list for the specific tag
	found := false

	for !found {
		for _, image := range list {
			for _, imageTag := range image.RepoTags {
				if tag == imageTag {
					found = true
					break
				}
			}
		}
	}

	return found, nil
}
