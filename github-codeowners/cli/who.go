package cli

import (
	"errors"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/codeowners"
	"github.com/spf13/cobra"
)

func WhoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "who",
		Short:   "lists the owner of the specified file",
		Example: "github-codewners who README.md",
		Args:    cobra.ExactArgs(1),
		RunE:    runWho,
	}

	cmd.Flags().BoolP("rule", "r", false, "print the rule which the file matched against")
	cmd.Flags().StringP("output", "o", "simple", "how to output format eg: simple, jsonl, csv")

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

	output, err := internal.GetOutput(cmd)
	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(coFilePath)
	if err != nil {
		return err
	}

	result := co.CalcOwnership(args[0])

	switch output {
	case "simple":
		internal.PrintSimple(cmd, []codeowners.CalcResult{result}, internal.PrintOpts{Path: false, Owners: !printRule,  Rule: printRule})
	case "csv":
		internal.PrintCsv(cmd, []codeowners.CalcResult{result}, internal.PrintOpts{Path: true, Owners: true, Rule: printRule})
	case "jsonl":
		internal.PrintJsonl(cmd, []codeowners.CalcResult{result})
	default:
		return errors.New("output type not implemented")
	}

	return nil
}
