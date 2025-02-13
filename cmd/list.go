package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all RFCs",
	Run: func(cmd *cobra.Command, args []string) {
		rfcDir := "rfc"
		files, err := os.ReadDir(rfcDir)
		if err != nil {
			fmt.Println("Error reading RFC directory:", err)
			return
		}

		fmt.Println("Existing RFCs:")
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".md" {
				fmt.Println("- " + file.Name())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
