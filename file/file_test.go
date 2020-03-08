package file

import (
	"io"
	"os"
	"testing"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Open(name string) (file, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

func TestTemplateFile_Replace(t *testing.T) {

}

func Test_replaceTemplates(t *testing.T) {
	replacements := make(map[string]string)
	replacements["one"] = "1"
	replacements["two"] = "2"

	content := "one should be 1 and two should be 2"
	expectedContent := "1 should be 1 and 2 should be 2"

	newContent := replaceTemplates(content, replacements)

	if expectedContent != newContent {
		t.Log("expected", expectedContent, " got ", newContent)
		t.Fail()
		return
	}
}
