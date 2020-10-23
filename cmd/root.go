package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ht",
	Short: "yama",
	Long: `The project aims to make an open source, multi-language, local development environment
that prioritizes role-based startup, service dependencies, testing and config generation
while remaining useable and unopinionated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
