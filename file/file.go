package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	templateString = "myapp"
)

// TemplateFile struct to operate of Templated Files
type TemplateFile struct {
	Name string
}

// Replace replaces templated files for the correct value
func (template TemplateFile) Replace() error {
	return filepath.Walk(template.Name, template.replace)
}

func (template TemplateFile) replace(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*", fi.Name())

	if err != nil {
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), templateString, template.Name, -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}
