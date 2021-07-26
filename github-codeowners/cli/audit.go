package cli

import (
	"errors"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal/flags"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal/outputs"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/codeowners"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/files"
	"github.com/spf13/cobra"
)

func AuditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "audit",
		Short:   "list the owners for all files",
		Example: "github-codewners audit",
		RunE:    runAudit,
	}

	cmd.Flags().BoolP("rule", "r", false, "print the rule which the file matched against")
	cmd.Flags().StringP("output", "o", "simple", "how to output format eg: simple, jsonl, csv")

	return cmd
}

func runAudit(cmd *cobra.Command, args []string) error {
	dir, err := flags.GetTrimmedFlag(cmd, "dir")
	if err != nil {
		return err
	}

	coFilePath, err := flags.GetCodeOwnersFilePath(cmd)
	if err != nil {
		return err
	}

	output, err := flags.GetOutput(cmd)
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

	paths, err := files.FindRecursively(dir)

	result := co.CalcManyOwnerships(paths)

	switch output {
	case "simple":
		outputs.PrintSimple(cmd, result, outputs.PrintOpts{Path: true, Owners: !printRule,  Rule: printRule})
	case "csv":
		outputs.PrintCsv(cmd, result, outputs.PrintOpts{Path: true, Owners: true, Rule: printRule})
	case "jsonl":
		err := outputs.PrintJsonl(cmd, result)
		if err != nil{
			return err
		}

	default:
		return errors.New("output type not implemented")
	}

	return nil
}
