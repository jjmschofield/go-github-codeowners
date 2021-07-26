package codeowners

import (
	"bufio"
	"errors"
	"github.com/sabhiram/go-gitignore"
	"os"
	"regexp"
	"strings"
)

type Codeowners struct {
	rules []coRule
}

type coRule struct {
	line    string
	owners  []string
	matcher *ignore.GitIgnore
}

type CalcResult struct {
	Owners []string
	Rule   string
}

func (c *Codeowners) CalcOwnership(path string) CalcResult {
	for i := 0; i < len(c.rules); i++ {
		rule := c.rules[i]

		if rule.matcher.MatchesPath(path) {
			return CalcResult{
				Owners: rule.owners,
				Rule:   rule.line,
			}
		}
	}

	return CalcResult{}
}

func FromFile(path string) (*Codeowners, error) {
	rules, err := readRulesFromFile(path)

	if err != nil {
		return nil, err
	}

	// We reverse the matchers so that the first matching rule encountered
	// will be the last from CODEOWNERS, respecting precedence correctly and performantly
	reversed := reverse(rules)

	return &Codeowners{
		rules: reversed,
	}, nil
}

func readRulesFromFile(path string) ([]coRule, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var ownerRegex = regexp.MustCompile(`(^@[a-zA-Z0-9_\-/]*$)|(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)

	var rules []coRule

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()

		if len(line) < 1 || line[0:1] == "#" {
			continue
		}

		parts := strings.Split(line, " ")

		matcher := ignore.CompileIgnoreLines(parts[0])

		var owners []string

		for i := 1; i < len(parts); i++ {
			clean := strings.TrimSpace(parts[i])

			if len(clean) > 0 {
				owners = append(owners, clean)
			}
		}

		for i := 0; i < len(owners); i++ {
			if !ownerRegex.MatchString(owners[i]) {
				return nil, errors.New("invalid owner syntax:" + owners[i])
			}
		}

		rules = append(
			rules,
			coRule{line, owners, matcher},
		)
	}

	return rules, scanner.Err()
}

func reverse(arr []coRule) []coRule {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}

	return arr
}
