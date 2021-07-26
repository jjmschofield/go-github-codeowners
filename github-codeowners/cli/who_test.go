package cli

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"testing"
)

func Test_Who_Help(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "--help"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Default_Codeowners(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "README.md"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Default_Codeowners_Print_Rule(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "-r", "README.md"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Selected_Codeowners(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "-c fixtures/CODEOWNERS_REFERENCE", "some-file.js"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Selected_Codeowners_Print_Rule(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "-c fixtures/CODEOWNERS_REFERENCE", "-r", "/docs/some-doc.txt"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Dir(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "-d ../../", "some-file.js"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Who_Dir_Selected_Codeowners(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"who", "-d fixtures", "-c CODEOWNERS_REFERENCE", "some-file.js"}...)
	cupaloy.SnapshotT(t, out)
}