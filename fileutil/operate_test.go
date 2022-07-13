package fileutil_test

import (
	"github.com/bychannel/goutils/fileutil"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	file, err := fileutil.CreateFile("./testdata/test.txt", 0664, 0666)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.Remove(file.Name()))
	}

	file, err = fileutil.CreateFile("./testdata/sub/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}

	file, err = fileutil.CreateFile("./testdata/sub/sub2/test.txt", 0664, 0777)
	if assert.NoError(t, err) {
		assert.Equal(t, "./testdata/sub/sub2/test.txt", file.Name())
		assert.NoError(t, file.Close())
		assert.NoError(t, os.RemoveAll("./testdata/sub"))
	}
}

func TestUnzip(t *testing.T) {
	assert.Error(t, fileutil.Unzip("/path-not-exists", ""))
}
