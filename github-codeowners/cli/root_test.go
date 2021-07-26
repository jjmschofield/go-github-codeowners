package cli

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"testing"
)

func Test_Root(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), "")
	cupaloy.SnapshotT(t, out)
}
