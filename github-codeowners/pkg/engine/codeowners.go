package engine

import (
	"bufio"
	"os"
)

type Codeowners struct {
	lines []string
}

func FromFile(path string) (*Codeowners, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()

		if line[0:1] == "#" {
			continue
		}

		lines = append(lines, scanner.Text())
	}

	return &Codeowners{ lines }, scanner.Err()
}
