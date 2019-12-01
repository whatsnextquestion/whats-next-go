package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(booksCmd)
}

var booksCmd = &cobra.Command{
	Use:   "books",
	Short: "Shows the books library",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("some books")
	},
}
