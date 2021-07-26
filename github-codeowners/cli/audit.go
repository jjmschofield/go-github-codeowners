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

type auditOpts struct {
	dir       string
	coPath    string
	output    string
	printRule bool
}

func runAudit(cmd *cobra.Command, args []string) error {
	opts, err := getAuditOpts(cmd)
	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(opts.coPath)
	if err != nil {
		return err
	}

	paths, err := files.FindRecursively(opts.dir)
	if err != nil {
		return err
	}

	result := co.CalcManyOwnerships(paths)

	switch opts.output {
	case "simple":
		outputs.PrintSimple(
			cmd,
			result,
			outputs.PrintOpts{Path: true, Owners: !opts.printRule, Rule: opts.printRule},
		)
	case "csv":
		outputs.PrintCsv(
			cmd,
			result,
			outputs.PrintOpts{Path: true, Owners: true, Rule: opts.printRule},
		)
	case "jsonl":
		err := outputs.PrintJsonl(cmd, result)
		if err != nil {
			return err
		}

	default:
		return errors.New("output type not implemented")
	}

	return nil
}

func getAuditOpts(cmd *cobra.Command) (auditOpts, error) {
	dir, err := flags.GetTrimmedFlag(cmd, "dir")
	if err != nil {
		return auditOpts{}, err
	}

	coPath, err := flags.GetCodeOwnersFilePath(cmd)
	if err != nil {
		return auditOpts{}, err
	}

	output, err := flags.GetOutput(cmd)
	if err != nil {
		return auditOpts{}, err
	}

	printRule, err := cmd.Flags().GetBool("rule")
	if err != nil {
		return auditOpts{}, err
	}

	return auditOpts{
		dir,
		coPath,
		output,
		printRule,
	}, nil
}
