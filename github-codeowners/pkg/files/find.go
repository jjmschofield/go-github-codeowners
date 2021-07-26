package files

import (
	ignore "github.com/sabhiram/go-gitignore"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func FindRecursively(path string) ([]string, error) {
	return walk(path, "", []foundIgnore{})
}

type foundIgnore struct {
	dir    string
	ignore *ignore.GitIgnore
}

func walk(root string, dir string, ignores []foundIgnore) ([]string, error) {
	dirAbs := GetAbsPath(filepath.Join(root, dir))

	if _, err := os.Stat(filepath.Join(dirAbs, ".gitignore")); err == nil {
		newIgnore, err := ignore.CompileIgnoreFile(filepath.Join(dirAbs, ".gitignore"))

		if err != nil {
			return nil, err
		}

		ignores = append(ignores, foundIgnore{
			dir:    dir,
			ignore: newIgnore,
		})
	}

	files, err := ioutil.ReadDir(dirAbs)

	if err != nil {
		return nil, err
	}

	var found []string

	for i := 0; i < len(files); i++ {
		file := files[i]
		relativePath := filepath.Join(dir, file.Name())

		if isIgnored(file, dir, ignores) {
			continue
		}

		if file.IsDir() && file.Name() != ".git" {
			children, err := walk(root, relativePath, ignores)

			if err != nil {
				return nil, err
			}

			found = append(found, children...)
		} else {
			found = append(found, relativePath)
		}
	}

	return found, nil
}

func isIgnored(file os.FileInfo, dir string, ignores []foundIgnore) bool{
	if file.Name() == ".gitignore" {
		return true
	}

	if file.IsDir() && file.Name() == ".git"{
		return true
	}

	relativePath := filepath.Join(dir, file.Name())

	for i := 0; i < len(ignores); i++ {
		thisIgnore := ignores[i]

		// We strip alter the file path so that it becomes relative to the ignore to be tried
		// If we don't do this, relative rules inside the ignore file will not be respected
		ignoreRelativePath := strings.Replace(relativePath, thisIgnore.dir, "", 1)

		if thisIgnore.ignore.MatchesPath(ignoreRelativePath) {
			return true
		}
	}

	return false
}