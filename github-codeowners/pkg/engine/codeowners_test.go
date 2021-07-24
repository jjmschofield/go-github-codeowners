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
	return filepath.Join(wd, path)
}

func TestCodeowners_FromFile_Should_Load_Rules_When_Valid(t *testing.T){
	codeowners, err := FromFile(getAbsPath("/fixtures/VALID_RULE"))

	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))

	rule := codeowners.rules[0]
	assert.Equal(t, 2, len(rule.owners))
	assert.Equal(t, []string{"@jjmschofield", "@jjmschofield2"}, rule.owners)
	assert.True(t, rule.matcher.MatchesPath("README.md"))
	assert.False(t, rule.matcher.MatchesPath("NOT_README.md"))
}


func TestCodeowners_FromFile_Should_Ignore_Comments(t *testing.T) {
	codeowners, err := FromFile(getAbsPath("/fixtures/COMMENTED_LINE"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))
	assert.Equal(t, "include @jjmschofield", codeowners.rules[0].line)
}

func TestCodeowners_FromFile_Should_Error_When_Owner_Is_Invalid(t *testing.T) {
	codeowners, err := FromFile(getAbsPath("/fixtures/INVALID_CODEOWNER"))
	assert.Nil(t, codeowners)
	assert.NotNil(t, err)
}
