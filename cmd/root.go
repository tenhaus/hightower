package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ht",
	Short: "yama",
	Long:  "Hightower - something something something",
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
