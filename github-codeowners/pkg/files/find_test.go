package files

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiles_findRecursively_Should_Find_All_Files(t *testing.T) {
	filePaths, err := FindRecursively("/fixtures/NO_IGNORE")

	assert.Nil(t, err)
	assert.NotNil(t, filePaths)

	assert.Contains(t, filePaths, "file-one")
	assert.Contains(t, filePaths, "level-1/file-two")
	assert.Contains(t, filePaths, "level-1/level-2/file-three")
	assert.Len(t, filePaths, 3)
}

func TestFiles_findRecursively_Should_Ignore_Files_When_Root_Ignore(t *testing.T) {
	err := OverwriteFile(GetAbsPath("fixtures/ROOT_IGNORE/.gitignore"), `/ignored-one
/level-1/ignored-two
/level-1/level-2/ignored-three
ignored-four`)

	assert.Nil(t, err)

	filePaths, err := FindRecursively("/fixtures/ROOT_IGNORE")

	assert.Nil(t, err)
	assert.NotNil(t, filePaths)

	assert.NotContains(t, filePaths, "ignored-one", "ignored-one wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/ignored-two", "level-1/ignored-two wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/level-2/ignored-three", "level-1/level-2/ignored-three wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/level-2/ignored-four", "level-1/level-2/ignored-four wasn't ignored")
	assert.Len(t, filePaths, 3)
}

func TestFiles_findRecursively_Should_Ignore_Files_When_Sub_Ignore(t *testing.T) {
	err := OverwriteFile(GetAbsPath("fixtures/SUB_IGNORE/.gitignore"), `/ignored-one`)
	err = OverwriteFile(GetAbsPath("fixtures/SUB_IGNORE/level-1/.gitignore"), `/ignored-two
/level-2/ignored-three`)
	err = OverwriteFile(GetAbsPath("fixtures/SUB_IGNORE/level-1/level-2/.gitignore"), `ignored-four`)

	filePaths, err := FindRecursively("/fixtures/SUB_IGNORE")

	assert.Nil(t, err)
	assert.NotNil(t, filePaths)

	assert.NotContains(t, filePaths, "ignored-one", "ignored-one wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/ignored-two", "level-1/ignored-two wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/level-2/ignored-three", "level-1/level-2/ignored-three wasn't ignored")
	assert.NotContains(t, filePaths, "level-1/level-2/ignored-four", "level-1/level-2/ignored-four wasn't ignored")
	assert.Len(t, filePaths, 3)
}
