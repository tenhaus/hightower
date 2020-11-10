package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

// RunOptions contains parameters for Run
type RunOptions struct {
	Image string
	Pull  bool
}

// Run starts a service
func Run(ctx *context.Context, options RunOptions) {

	cli, err := NewClient()
	if err != nil {
		panic(err)
	}

	if options.Pull {
		reader, err := cli.ImagePull(*ctx, options.Image, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}

		io.Copy(os.Stdout, reader)
	}

	resp, err := cli.ContainerCreate(*ctx, &container.Config{
		Image: options.Image,
		// Cmd:   []string{"echo", "hello world"},
		Tty: false,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(*ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(*ctx, resp.ID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(*ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
