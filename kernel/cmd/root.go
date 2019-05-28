package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "gominix",
	Short: "gominix is interpreter for minix binary",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
