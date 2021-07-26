package internal

import (
	"github.com/spf13/cobra"
	"path"
	"strings"
)

func GetCodeOwnersFilePath(cmd *cobra.Command) (string, error) {
	rootPath, err := cmd.Flags().GetString("dir")
	coPath, err := cmd.Flags().GetString("codeowners")

	if err != nil {
		return "", err
	}

	result := path.Join(strings.TrimSpace(rootPath), strings.TrimSpace(coPath))

	return result, nil
}
