package runnablereadme

import (
	"fmt"

	runnablereadme "github.com/R3DRUN3/runnable-readme/pkg/runnable-readme"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:     "exec",
	Aliases: []string{"exec"},
	Short:   "Read a markdown",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := runnablereadme.ExecuteMarkdown(args[0])
		fmt.Println("Operation OK:", res)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
