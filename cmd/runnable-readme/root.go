package runnablereadme

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rure",
	Short: "RUnnable REadme - a simple CLI to execute shell instruction in markdown files",
	Long: `RUnnable REadme is a super fancy CLI (kidding)
One can use RUnnable REadme to execute console or sh instructions found on markdown file...documentation as code!`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI, sorry for that 😔  '%s'", err)
		os.Exit(1)
	}
}
