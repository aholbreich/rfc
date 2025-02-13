package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rfc",
	Short: "RFC management tool",
	Long:  `A CLI tool for managing RFCs inside a project repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to RFC tool! Use 'rfc --help' for commands.")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
