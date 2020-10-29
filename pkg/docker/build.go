package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

// BuildContext creates a docker context
func BuildContext(buf *bytes.Buffer, dir string) error {
	tw := tar.NewWriter(buf)

	filepath.Walk(dir, func(file string, fi os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(file)
		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil

	})

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}

	return nil
}

func Build() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	BuildContext(&buf, "test/base")

	options := types.ImageBuildOptions{
		Context:    &buf,
		Dockerfile: "test/base/Dockerfile",
		Tags:       []string{"chris:test"},
		NoCache:    true,
	}

	resp, err := cli.ImageBuild(ctx, &buf, options)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stderr, termFd, isTerm, nil)

}
