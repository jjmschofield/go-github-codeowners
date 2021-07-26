package cli

import (
	"github.com/bradleyjkemp/cupaloy"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli/internal"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Audit_Help(t *testing.T) {
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "--help"}...)
	cupaloy.SnapshotT(t, out)
}

func writeIgnores(t *testing.T) {
	err := files.OverwriteFile(files.GetAbsPath(".fixtures/COMPLEX_PROJECT/.gitignore"), `node_modules
explicit-ignore.js
overridden-ignore.js
override.txt
`)

	err = files.OverwriteFile(files.GetAbsPath(".fixtures/COMPLEX_PROJECT/deep/nested-ignore/.gitignore"), `!overridden-ignore.js
ignored-by-nested-rule.txt`)

	assert.Nil(t, err)
}

func Test_Audit(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Audit_Print_Rule(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT", "-r"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Audit_Csv(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT", "-o csv"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Audit_Jsonl(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT", "-o jsonl"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Audit_Csv_Print_Rule(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT", "-o csv", "-r"}...)
	cupaloy.SnapshotT(t, out)
}

func Test_Audit_Selected_Codeowners(t *testing.T) {
	writeIgnores(t)
	_, out, _ := internal.ExecuteCommand(RootCmd(), []string{"audit", "-d .fixtures/COMPLEX_PROJECT", "-c .fixtures/SIMPLE"}...)
	cupaloy.SnapshotT(t, out)
}

