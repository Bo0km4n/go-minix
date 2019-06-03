package cmd

import (
	"log"

	"github.com/Bo0km4n/go-minix/kernel/pkg/core"
	"github.com/spf13/cobra"
)

var (
	Input string
	Trace bool
)

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&Input, "input", "i", "", "Select file path of Minix binary")
	execCmd.Flags().BoolVarP(&Trace, "trace", "t", false, "Trace memory state")
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute the minix binary",
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.Boot(Input, Trace); err != nil {
			log.Fatal(err)
		}
	},
}
