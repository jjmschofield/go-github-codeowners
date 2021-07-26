package codeowners

import (
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeowners_calcOwnership_Should_Respect_Reference(t *testing.T) {
	// see https://docs.github.com/en/enterprise-server@2.21/github/creating-cloning-and-archiving-repositories/creating-a-repository-on-github/about-code-owners
	codeowners, err := FromFile(files.GetAbsPath("/fixtures/REFERENCE"))

	assert.Nil(t, err)

	assert.Equal(t, []string{"@global-owner1", "@global-owner2"}, codeowners.calcOwnership("some-random-file"))
	assert.Equal(t, []string{"@js-owner"}, codeowners.calcOwnership("nested/file.js"))
	assert.Equal(t, []string{"docs@example.com"}, codeowners.calcOwnership("docs.go"))
	assert.Equal(t, []string{"@doctocat"}, codeowners.calcOwnership("/build/logs/log"))
	assert.Equal(t, []string{"docs@example.com"}, codeowners.calcOwnership("anywhere/docs/getting-started.md")) // anywhere/ prevents /docs/ @doctocat from winning
	// assert.Equal(t, []string{"@global-owner1", "@global-owner2"}, codeowners.calcOwnership("anywhere/docs/build-app/troubleshooting.md")) // TODO - the underlying ignore lib does not respect this
	assert.Equal(t, []string{"@octocat"}, codeowners.calcOwnership("anywhere/apps/file"))
	assert.Equal(t, []string{"@doctocat"}, codeowners.calcOwnership("/docs/file"))
	assert.Equal(t, []string{"@doctocat"}, codeowners.calcOwnership("/docs/nested/file"))
	assert.Equal(t, []string{"@octocat"}, codeowners.calcOwnership("/apps/file"))
	assert.Equal(t, []string(nil), codeowners.calcOwnership("/apps/github"))
}

func TestCodeowners_FromFile_Should_Load_Rules_When_Valid(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/fixtures/VALID_RULE"))

	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))

	rule := codeowners.rules[0]
	assert.Equal(t, 2, len(rule.owners))
	assert.Equal(t, []string{"@jjmschofield", "@jjmschofield2"}, rule.owners)
	assert.True(t, rule.matcher.MatchesPath("README.md"))
	assert.False(t, rule.matcher.MatchesPath("NOT_README.md"))
}

func TestCodeowners_FromFile_Should_Reverse_Rules(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/fixtures/REVERSE"))

	assert.Nil(t, err)
	assert.Equal(t, 2, len(codeowners.rules))

	assert.Equal(t, "file-two", codeowners.rules[0].line)
	assert.Equal(t, "file-one", codeowners.rules[1].line)
}

func TestCodeowners_FromFile_Should_Ignore_Comments(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/fixtures/COMMENTED_LINE"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))
	assert.Equal(t, "include @jjmschofield", codeowners.rules[0].line)
}

func TestCodeowners_FromFile_Should_Error_When_Owner_Is_Invalid(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/fixtures/INVALID_CODEOWNER"))
	assert.Nil(t, codeowners)
	assert.NotNil(t, err)
}
