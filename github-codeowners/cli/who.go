package cli

import (
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
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
		RunE:    runWho,
	}

	cmd.Flags().BoolP("rule", "r", false, "Print the rule which the file matched against")

	return cmd
}

func runWho(cmd *cobra.Command, args []string) error {
	coFilePath, err := internal.GetCodeOwnersFilePath(cmd)
	if err != nil {
		return err
	}

	printRule, err := cmd.Flags().GetBool("rule")
	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(coFilePath)
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
