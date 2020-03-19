package file

import (
	b64 "encoding/base64"
	"io/ioutil"

	"github.com/elko-dev/spawn/platform"
)

// Reader to read files from file system
type Reader struct {
}

// AsBase64String reads file from path as a string and returns Base64 String
func (reader Reader) AsBase64String(fileName string) (string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(file), nil
}

// NewReader init
func NewReader() platform.Secrets {
	return Reader{}
}
