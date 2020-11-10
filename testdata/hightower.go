package main

import (
	"github.com/tenhaus/hightower/pkg/api"
)

var postgres = api.DockerService{
	Name:       "Postgres",
	Dockerfile: "testdata/dockerproject/postgres/Dockerfile",
	Context:    "testdata/dockerproject/postgres",
}

func main() {
	postgres.Build()
	postgres.Run()
}

// Next steps

// * Build and deploy a docker service
// * Watch for changes and reload
//
// * Connect two services
// * Port forward
//
// * Up check
// * Post-up init
