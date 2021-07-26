package cli

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"testing"
)

func Test_VersionNumber(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(rootCmd, "version")
	cupaloy.SnapshotT(t, out)
}
