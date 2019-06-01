package cmd

import (
	"log"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core"
	"github.com/spf13/cobra"
)

var Input string

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&Input, "input", "i", "", "Select file path of Minix binary")
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute the minix binary",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Boot(Input); err != nil {
			log.Fatal(err)
		}
	},
}
