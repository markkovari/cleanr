package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cleanr",
	Short: "Cleanr is taking care of the temporarly created projects or flies for you",
	Long: `Most of the times you are just showcase something quickly and forgot about that
          project/file/whatever it is. Instead of nying a new computer you should keep your
          projects tidy. Cleanr helps.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cleanr")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
