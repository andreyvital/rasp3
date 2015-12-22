package fs_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/CentaurWarchief/rasp3/fs"
	"github.com/stretchr/testify/assert"
)

func TestReadEmptyDirectory(t *testing.T) {
	var dir string
	var err error

	if dir, err = ioutil.TempDir(os.TempDir(), "TestReadEmptyDirectory"); err != nil {
		t.Fail()
		return
	}

	defer os.Remove(dir)

	assert.Len(t, fs.Readdir(dir), 0)
}

func TestReaddir(t *testing.T) {
	var dir string
	var err error

	if dir, err = ioutil.TempDir(os.TempDir(), "TestReaddir"); err != nil {
		t.Fail()
		return
	}

	defer os.RemoveAll(dir)

	ioutil.TempFile(dir, "1")
	ioutil.TempFile(dir, "2")
	ioutil.TempFile(dir, "3")

	assert.Len(t, fs.Readdir(dir), 3)
}
