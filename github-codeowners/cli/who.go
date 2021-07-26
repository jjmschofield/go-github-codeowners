package cli

import (
	"errors"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal/flags"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal/outputs"
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

type whoOpts struct {
	coPath    string
	output    string
	printRule bool
}

func runWho(cmd *cobra.Command, args []string) error {
	opts, err := getWhoOpts(cmd)
	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(opts.coPath)
	if err != nil {
		return err
	}

	result := co.CalcOwnership(args[0])

	switch opts.output {
	case "simple":
		outputs.PrintSimple(
			cmd,
			[]codeowners.CalcResult{result},
			outputs.PrintOpts{Path: false, Owners: !opts.printRule, Rule: opts.printRule},
		)
	case "csv":
		outputs.PrintCsv(
			cmd,
			[]codeowners.CalcResult{result},
			outputs.PrintOpts{Path: true, Owners: true, Rule: opts.printRule},
		)
	case "jsonl":
		err := outputs.PrintJsonl(cmd, []codeowners.CalcResult{result})
		if err != nil {
			return err
		}
	default:
		return errors.New("output type not implemented")
	}

	return nil
}

func getWhoOpts(cmd *cobra.Command) (whoOpts, error) {
	coPath, err := flags.GetCodeOwnersFilePath(cmd)
	if err != nil {
		return whoOpts{}, err
	}

	output, err := flags.GetOutput(cmd)
	if err != nil {
		return whoOpts{}, err
	}

	printRule, err := cmd.Flags().GetBool("rule")
	if err != nil {
		return whoOpts{}, err
	}

	return whoOpts{
		coPath,
		output,
		printRule,
	}, nil
}
