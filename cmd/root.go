package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tenhaus/hightower/pkg/config"
	"github.com/tenhaus/hightower/pkg/docker"
)

var rootCmd = &cobra.Command{
	Use:   "ht",
	Short: "yama",
	Long:  description,
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run the development environment",
	Run: func(cmd *cobra.Command, args []string) {

		// Grab our config
		config, err := config.Parse()
		if err != nil {
			log.Fatalf("There was an error loading the configuration: %v", err)

		}

		// Make sure the entrypoint exists
		_, err = os.Stat(config.EntryPoint)
		if os.IsNotExist(err) {
			log.Fatalf("Entrypoint (%v) does not exist", config.EntryPoint)

		}

		// Grab an image from the cli
		hightowerCmd := exec.Command("go", "run", config.EntryPoint)
		if err := hightowerCmd.Run(); err != nil {
			log.Fatalf("There was an error executing %v: %v", config.EntryPoint, err)
		}
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the environment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		docker.Build(&ctx, docker.BuildOptions{
			Path:       "test/base",
			Dockerfile: "test/base/Dockerfile",
			Tag:        "chris:test",
			Cache:      false,
		})
	},
}

// Execute is the Cobra entrypoint of our application
func Execute() {
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(buildCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const description = `The project aims to make an open source, multi-language, local development environment
that prioritizes role-based startup, service dependencies, testing and config generation
while remaining useable and unopinionated.`
