package fs

import (
	"io/ioutil"
)

//ReadFileString read al content from file and return as string
func ReadFileString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}
