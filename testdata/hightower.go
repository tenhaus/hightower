package main

import (
	"context"

	"github.com/tenhaus/hightower/pkg/docker"
)

// DockerContainer is a first pass at a running service
// and will later be wrapped in an interface
// and thrown in api/
type DockerContainer struct {
	Dockerfile string
	Context    string
}

// DockerService ...
type DockerService interface {
	Start()
	Build()
}

// DockerServiceRunner ...
type DockerServiceRunner struct {
	Container DockerContainer
}

func (r *DockerServiceRunner) Start() {
}

func (r *DockerServiceRunner) Build() {
	ctx := context.Background()
	docker.Build(&ctx, docker.BuildOptions{
		DisplayName: "Test Service",
		Dockerfile:  r.Container.Dockerfile,
		Path:        r.Container.Context,
		Cache:       false,
	})
}

func main() {
	container := DockerContainer{
		Dockerfile: "testdata/dockerservice/Dockerfile",
		Context:    "testdata/dockerservice",
	}

	runner := DockerServiceRunner{Container: container}
	runner.Build()

	// Next steps

	// * Build and deploy a docker service
	// * Watch for changes and reload
	//
	// * Connect two services
	// * Port forward
	//
	// * Up check
	// * Post-up init

}
