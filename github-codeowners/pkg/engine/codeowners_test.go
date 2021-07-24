package engine

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func getAbsPath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, path);
}

func TestCodeowners_FromFile_CommentedOutLines(t *testing.T) {
	codeowners, err := FromFile(getAbsPath("/fixtures/COMMENTED_LINE"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.lines))
	assert.Equal(t, "include", codeowners.lines[0])
}
