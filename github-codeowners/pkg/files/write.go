package files

import (
	"io/ioutil"
)

func OverwriteFile(file string, contents string) error {
	return ioutil.WriteFile(file, []byte(contents), 0644)
}
