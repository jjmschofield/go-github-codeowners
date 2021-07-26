package main

import (
	"fmt"
	"github.com/jjmschofield/go-github-codeowners/github-codeowners/cli"
	"os"
)

func main() {
	root := cli.RootCmd()

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
