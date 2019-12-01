package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addBookCmd)
}

var addBookCmd = &cobra.Command{
	Use:   "add-book",
	Short: "Adds a book from Project Gutenberg",
	Long:  `Add a book by Project Gutenberg link such as https://www.gutenberg.org/ebooks/<book id>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adding book")
	},
}

func init() {
	addBookCmd.PersistentFlags().StringP("url", "u", "", "url in format https://www.gutenberg.org/ebooks/<book id> ")
}
