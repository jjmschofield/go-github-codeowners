package cli

import (
	"errors"
	"github.com/spf13/cobra"
)

func RootCmd () *cobra.Command{
	cmd := &cobra.Command{
		Use:   "github-codeowners",
		Short: "A collection of tools to make the most out of github CODEOWNER files",
		RunE:  runRoot,
	}

	addFlags(cmd)

	addSubCommands(cmd)

	return cmd
}

func addFlags (cmd *cobra.Command){
	cmd.PersistentFlags().StringP(
		"codeowners",
		"c",
		".github/CODEOWNERS",
		"Path to the CODEOWNERS file relative to your current directory",
	)

	cmd.PersistentFlags().StringP(
		"dir",
		"d",
		"./",
		"Path to the directory to explore with github-codeowners",
	)
}

func addSubCommands (cmd *cobra.Command){
	cmd.AddCommand(VersionCmd())
	cmd.AddCommand(WhoCmd())
	cmd.AddCommand(AuditCmd())
}

func runRoot(cmd *cobra.Command, args []string) error {
	return errors.New("specify a command to run")
}