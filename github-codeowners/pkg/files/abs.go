package files

import (
	"os"
	"path/filepath"
)

func GetAbsPath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, path)
}