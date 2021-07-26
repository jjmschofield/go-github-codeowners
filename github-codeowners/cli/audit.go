package cli

import (
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/codeowners"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/files"
	"github.com/spf13/cobra"
	"strings"
)

func AuditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "audit",
		Short:   "list the owners for all files",
		Example: "github-codewners audit",
		RunE:    runAudit,
	}

	return cmd
}

func runAudit(cmd *cobra.Command, args []string) error {
	dir, err := internal.GetTrimmedFlag(cmd, "dir")
	if err != nil {
		return err
	}

	coFilePath, err := internal.GetCodeOwnersFilePath(cmd)
	if err != nil {
		return err
	}

	co, err := codeowners.FromFile(coFilePath)
	if err != nil {
		return err
	}

	paths, err := files.FindRecursively(dir)

	result := co.CalcManyOwnerships(paths)

	for i:=0; i< len(result); i++{
		cmd.Println(result[i].Path + " " + strings.Join(result[i].Owners, " ") )
	}

	return nil
}
