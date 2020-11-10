package api

import (
	"context"

	"github.com/tenhaus/hightower/pkg/docker"
)

// DockerService represents a Docker service
type DockerService struct {
	Name       string
	Dockerfile string
	Context    string
	image      string
}

// Build builds a docker service
func (s *DockerService) Build() {
	ctx := context.Background()
	options := docker.BuildOptions{
		Dockerfile:  s.Dockerfile,
		DisplayName: s.Name,
		Path:        s.Context,
		Tag:         "hightower:testimage",
	}

	docker.Build(&ctx, options)
	s.image = "hightower:testimage"
}

// Run starts a docker service
func (s *DockerService) Run() {
	ctx := context.Background()
	options := docker.RunOptions{
		Image: s.image,
		Pull:  false,
	}

	docker.Run(&ctx, options)
}
