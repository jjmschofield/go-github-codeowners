package flags

import (
	"errors"
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



func GetOutput(cmd *cobra.Command) (string, error) {
	output, err := GetTrimmedFlag(cmd, "output")
	if err != nil {
		return "", err
	}

	if output != "simple" && output != "csv" && output != "jsonl" {
		return "", errors.New("output must be one of: simple, csv, jsonl")
	}

	return output, nil
}
