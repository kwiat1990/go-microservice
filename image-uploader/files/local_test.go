package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
	dir, err := ioutil.TempDir("", "files")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, 200)
	if err != nil {
		t.Fatal(err)
	}

	return l, dir, func() {
		// cleanup func
		// os.RemoveAll
	}
}

func TestSaveContentsOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello world"
	l, dir, cleanup := setupLocal(t)
	defer cleanup()

	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// check the file has been correctly written
	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	// check the contents of the file
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

func TestGetContentsAndWrite(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello world"
	l, _, cleanup := setupLocal(t)
	defer cleanup()

	// save a file
	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// read the file back
	r, err := l.Get(savePath)
	assert.NoError(t, err)
	defer r.Close()

	// read the full contents of the reader
	d, _ := ioutil.ReadAll(r)
	assert.Equal(t, fileContents, string(d))
}
