package codeowners

import (
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/pkg/files"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeowners_calcOwnership_Should_Respect_Reference(t *testing.T) {
	// see https://docs.github.com/en/enterprise-server@2.21/github/creating-cloning-and-archiving-repositories/creating-a-repository-on-github/about-code-owners
	codeowners, err := FromFile(files.GetAbsPath("/.fixtures/REFERENCE"))

	assert.Nil(t, err)

	assert.Equal(t, []string{"@global-owner1", "@global-owner2"}, codeowners.CalcOwnership("some-random-file").Owners)
	assert.Equal(t, "*       @global-owner1 @global-owner2", codeowners.CalcOwnership("some-random-file").Rule)

	assert.Equal(t, []string{"@js-owner"}, codeowners.CalcOwnership("nested/file.js").Owners)
	assert.Equal(t, "*.js    @js-owner", codeowners.CalcOwnership("nested/file.js").Rule)

	assert.Equal(t, []string{"docs@example.com"}, codeowners.CalcOwnership("docs.go").Owners)
	assert.Equal(t, "*.go docs@example.com", codeowners.CalcOwnership("docs.go").Rule)

	assert.Equal(t, []string{"@doctocat"}, codeowners.CalcOwnership("/build/logs/log").Owners)
	assert.Equal(t, "/build/logs/ @doctocat", codeowners.CalcOwnership("/build/logs/log").Rule)

	assert.Equal(t, []string{"docs@example.com"}, codeowners.CalcOwnership("anywhere/docs/getting-started.md").Owners) // anywhere/ prevents /docs/ @doctocat from winning
	assert.Equal(t, "docs/*  docs@example.com", codeowners.CalcOwnership("anywhere/docs/getting-started.md").Rule)

	// assert.Equal(t, []string{"@global-owner1", "@global-owner2"}, codeowners.CalcOwnership("anywhere/docs/build-app/troubleshooting.md")) // TODO - the underlying ignore lib does not respect this

	assert.Equal(t, []string{"@octocat"}, codeowners.CalcOwnership("anywhere/apps/file").Owners)
	assert.Equal(t, "apps/ @octocat", codeowners.CalcOwnership("anywhere/apps/file").Rule)

	assert.Equal(t, []string{"@doctocat"}, codeowners.CalcOwnership("/docs/file").Owners)
	assert.Equal(t, "/docs/ @doctocat", codeowners.CalcOwnership("/docs/file").Rule)

	assert.Equal(t, []string{"@doctocat"}, codeowners.CalcOwnership("/docs/nested/file").Owners)
	assert.Equal(t, "/docs/ @doctocat", codeowners.CalcOwnership("/docs/nested/file").Rule)

	assert.Equal(t, []string{"@octocat"}, codeowners.CalcOwnership("/apps/file").Owners)
	assert.Equal(t, "/apps/ @octocat", codeowners.CalcOwnership("/apps/file").Rule)

	assert.Equal(t, []string(nil), codeowners.CalcOwnership("/apps/github").Owners)
	assert.Equal(t, "/apps/github ", codeowners.CalcOwnership("/apps/github").Rule)
}

func TestCodeowners_FromFile_Should_Load_Rules_When_Valid(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/.fixtures/VALID_RULE"))

	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))

	rule := codeowners.rules[0]
	assert.Equal(t, 2, len(rule.owners))
	assert.Equal(t, []string{"@jjmschofield", "@jjmschofield2"}, rule.owners)
	assert.True(t, rule.matcher.MatchesPath("README.md"))
	assert.False(t, rule.matcher.MatchesPath("NOT_README.md"))
}

func TestCodeowners_FromFile_Should_Reverse_Rules(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/.fixtures/REVERSE"))

	assert.Nil(t, err)
	assert.Equal(t, 2, len(codeowners.rules))

	assert.Equal(t, "file-two", codeowners.rules[0].line)
	assert.Equal(t, "file-one", codeowners.rules[1].line)
}

func TestCodeowners_FromFile_Should_Ignore_Comments(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/.fixtures/COMMENTED_LINE"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(codeowners.rules))
	assert.Equal(t, "include @jjmschofield", codeowners.rules[0].line)
}

func TestCodeowners_FromFile_Should_Error_When_Owner_Is_Invalid(t *testing.T) {
	codeowners, err := FromFile(files.GetAbsPath("/.fixtures/INVALID_CODEOWNER"))
	assert.Nil(t, codeowners)
	assert.NotNil(t, err)
}
