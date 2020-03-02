package directory

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	templateString = "myapp"
	gitPath        = ".git"
)

// TemplateDirectory struct to operate on Directories
type TemplateDirectory struct {
	Name string
}

// Replace replaces templated files for the correct value
func (template TemplateDirectory) Replace() error {
	log.WithFields(log.Fields{
		"templateString":  templateString,
		"replacementName": template.Name,
	}).Debug("Starting replacement")

	return filepath.Walk(template.Name, template.replace)
}
func (template TemplateDirectory) replace(path string, f os.FileInfo, err error) error {
	contextLogger := log.WithFields(log.Fields{
		"path":            path,
		"templateString":  templateString,
		"replacementName": template.Name,
		"err":             err,
	})

	contextLogger.Debug("Starting replacement for path ", path)

	if strings.Contains(path, gitPath) {
		return nil
	}
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		return nil
	}

	if err != nil {
		return err
	}

	contextLogger.Debug("No error, continuing")
	if !f.IsDir() {
		contextLogger.Debug("Not a directory, continuing")
		return nil //
	}
	contextLogger.Debug("Is a directory, checking if name matches")

	if name := f.Name(); strings.Contains(name, templateString) {

		dir := filepath.Dir(path)
		contextLogger.Debug("Found directory ", dir)

		newname := strings.Replace(name, templateString, template.Name, 1)
		contextLogger.Debug("New Name ", newname)

		newpath := filepath.Join(dir, newname)
		contextLogger.Debug("New Path ", newpath)
		err := os.Rename(path, newpath)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
