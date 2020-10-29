package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/tenhaus/hightower/pkg/fs"
)

// CreateBuildContext creates a docker context
func CreateBuildContext(buf *bytes.Buffer, dir string) error {
	tw := tar.NewWriter(buf)
	files, err := fs.Walk(dir, true)

	if err != nil {
		return err
	}

	// Add each file to the tar
	for _, file := range files {
		header, err := tar.FileInfoHeader(file.Info, file.Path)

		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(file.Path)

		// Write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Copy files
		if !file.Info.IsDir() {
			data, err := os.Open(file.Path)

			if err != nil {
				return err
			}

			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	}

	// Write the tar
	if err := tw.Close(); err != nil {
		return err
	}

	return nil
}

// BuildOptions exposes a subset of docker's types.ImageBuildOptions
type BuildOptions struct {
	Path       string
	Dockerfile string
	Tag        string
	Cache      bool
}

// Build creates a docker image from the specified Dockerfile and context
func Build(ctx *context.Context, options BuildOptions) {

	// Store the image
	var buf bytes.Buffer

	// Initialize a docker client
	cli, err := NewClient()

	if err != nil {
		panic(err)
	}

	// Create a tar of the files needed to support
	// our Dockerfile instructions
	CreateBuildContext(&buf, options.Path)

	// Configure the build
	buildOptions := types.ImageBuildOptions{
		Context:    &buf,
		Dockerfile: options.Dockerfile,
		Tags:       []string{options.Tag},
		NoCache:    !options.Cache,
	}

	// Build
	resp, err := cli.ImageBuild(*ctx, &buf, buildOptions)

	if err != nil {
		panic(err)
	}

	// Make sure we close the reference to the builder
	defer resp.Body.Close()

	// Copy the client's build output to stdout
	termFd, isTerm := term.GetFdInfo(os.Stdout)
	jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stdout, termFd, isTerm, nil)
}
