package file

import (
	"io/ioutil"
)

// Reader to read files from file system
type Reader struct {
}

// AsString reads file from path as a string
func (reader Reader) AsString(fileName string) (string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
