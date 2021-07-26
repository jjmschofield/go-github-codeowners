package internal

import (
	"encoding/json"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/codeowners"
	"github.com/spf13/cobra"
	"strings"
)

type PrintOpts struct {
	Path   bool
	Rule   bool
	Owners bool
}

func PrintSimple(cmd *cobra.Command, results []codeowners.CalcResult, opts PrintOpts) {
	for i := 0; i < len(results); i++ {
		toPrint := resultToSlice(results[i], opts)
		cmd.Println(strings.Join(toPrint, "\t"))
	}
}

func PrintCsv(cmd *cobra.Command, results []codeowners.CalcResult, opts PrintOpts) {
	for i := 0; i < len(results); i++ {
		toPrint := resultToSlice(results[i], opts)
		cmd.Println(strings.Join(toPrint, ","))
	}
}

func PrintJsonl(cmd *cobra.Command, results []codeowners.CalcResult) error {
	for i := 0; i < len(results); i++ {
		jsonl, err := json.Marshal(results[i])
		if err != nil {
			return err
		}

		cmd.Println(string(jsonl))
	}
	return nil
}

func resultToSlice(result codeowners.CalcResult, opts PrintOpts) []string {
	var slice []string

	if opts.Path {
		slice = append(slice, result.Path)
	}

	if opts.Rule {
		slice = append(slice, result.Rule)
	}

	if opts.Owners {
		slice = append(slice, result.Owners...)
	}

	return slice
}
