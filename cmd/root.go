package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tenhaus/hightower/pkg/docker"
)

var rootCmd = &cobra.Command{
	Use:   "ht",
	Short: "yama",
	Long: `The project aims to make an open source, multi-language, local development environment
that prioritizes role-based startup, service dependencies, testing and config generation
while remaining useable and unopinionated.`,
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run the development environment",
	Run: func(cmd *cobra.Command, args []string) {
		docker.Run()
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
