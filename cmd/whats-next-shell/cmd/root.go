package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "whats-next-shell",
	Short: "What's Next gives answers to what's next in life, provided that the answer is in books. ",
	Long: `What's Next gives answers to what's next in life, provided that the answer is in books. 
Your questions need to be short, currently only up-to 5 words are supported (demo version).

You need to add some books from Project Gutenberg and keep rolling.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//},
}

// Execute is the root command for whats-next-shell
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// TODO: some viper logic here
}
