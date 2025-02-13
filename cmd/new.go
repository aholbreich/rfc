package cmd

import (
	"fmt"
	"unicode"

	"github.com/aholbreich/rfc-tool/internal/rfc"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [title].",
	Short: "Create a new RFC with title.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]

		// Validate the title (e.g., check if it's non-empty and meets other criteria)
		if title == "" {
			fmt.Println("Error: Title cannot be empty")
			cmd.Usage() // Show usage instructions
			return
		}

		if !isValidTitle(title) {
			fmt.Println("Error: Title contains invalid characters")
			return
		}

		rfc.New(title)

	},
}

func isValidTitle(title string) bool {

	if len(title) < 3 || len(title) > 1024 {
		return false
	}

	for _, char := range title {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != ' ' && char != '-' {
			return false
		}
	}
	return true
}

func init() {
	rootCmd.AddCommand(newCmd)
}
