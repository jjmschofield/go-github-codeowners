package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of github-codeowners",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("v0.0.1")
	},
}
