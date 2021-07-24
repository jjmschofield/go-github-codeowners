package cli

import (
"errors"
"fmt"
"github.com/spf13/cobra"
"os"
)

var rootCmd = &cobra.Command{
	Use:   "github-codeowners",
	Short: "A collection of tools to make the most out of github CODEOWNER files",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("specify a command to run")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
