package internal

import (
	"github.com/spf13/cobra"
	"path"
	"strings"
)

func GetTrimmedFlag(cmd *cobra.Command, name string) (string, error) {
	value, err := cmd.Flags().GetString(name)
	return strings.TrimSpace(value), err
}

func GetCodeOwnersFilePath(cmd *cobra.Command) (string, error) {
	coPath, err := GetTrimmedFlag(cmd, "codeowners")
	dirPath, err := GetTrimmedFlag(cmd, "dir")

	if err != nil {
		return "", err
	}

	if coPath == ".github/CODEOWNERS" {
		return path.Join(dirPath, coPath), nil
	}

	return coPath, nil
}
