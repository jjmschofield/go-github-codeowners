package cli

import (
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of github-codeowners",
		RunE:  runVersion,
	}

	return cmd
}

func runVersion(cmd *cobra.Command, args []string) error {
	cmd.Println("v0.0.1")
	return nil
}
