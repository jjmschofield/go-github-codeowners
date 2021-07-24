package cli

import (
	"bytes"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/spf13/cobra"
	"testing"
)

func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func Test_VersionNumber(t *testing.T) {
	_, out, _ := ExecuteCommandC(rootCmd, "version")
	cupaloy.SnapshotT(t, out)
}
