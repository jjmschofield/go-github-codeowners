package cli

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal/test"
	"testing"
)

func Test_VersionNumber(t *testing.T) {
	_, out, _ := test.ExecuteCommand(RootCmd(), "version")
	cupaloy.SnapshotT(t, out)
}
