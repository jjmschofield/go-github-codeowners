package cli

import (
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/codeowners"
	"github.com/spf13/cobra"
	"strings"
)

func WhoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "who",
		Short:   "Prints the owner of the specified file",
		Example: "github-codewners who README.md",
		Args:    cobra.ExactArgs(1),
		RunE: runWho,
	}

	cmd.Flags().BoolP("rule", "r", false, "Print the rule which the file matched against")

	return cmd
}

func runWho(cmd *cobra.Command, args []string) error {
	coPath, err := cmd.Flags().GetString("codeowners")
	printRule, err := cmd.Flags().GetBool("rule")

	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(strings.TrimSpace(coPath))

	if err != nil {
		return err
	}

	result := co.CalcOwnership(args[0])

	if !printRule {
		cmd.Println(strings.Join(result.Owners, " "))
	} else {
		cmd.Println(result.Rule)
	}

	return nil
}