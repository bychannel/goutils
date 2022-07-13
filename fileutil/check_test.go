package fileutil_test

import (
	"github.com/bychannel/goutils/fileutil"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

//goland:noinspection GoBoolExpressions
func TestCommon(t *testing.T) {
	assert.Equal(t, "", fileutil.FileExt("testdata/testjpg"))
	assert.Equal(t, ".jpg", fileutil.Suffix("testdata/test.jpg"))

	// IsZipFile
	assert.False(t, fileutil.IsZipFile("testdata/test.jpg"))
	assert.Equal(t, "test.jpg", fileutil.PathName("testdata/test.jpg"))

	assert.Equal(t, "test.jpg", fileutil.Name("path/to/test.jpg"))

	if runtime.GOOS == "windows" {
		assert.Equal(t, "path\\to", fileutil.Dir("path/to/test.jpg"))
	} else {
		assert.Equal(t, "path/to", fileutil.Dir("path/to/test.jpg"))
	}
}

func TestPathExists(t *testing.T) {
	assert.False(t, fileutil.PathExists(""))
	assert.False(t, fileutil.PathExists("/not-exist"))
	assert.True(t, fileutil.PathExists("testdata/test.jpg"))
}

func TestIsFile(t *testing.T) {
	assert.False(t, fileutil.IsFile(""))
	assert.False(t, fileutil.IsFile("/not-exist"))
	assert.True(t, fileutil.IsFile("testdata/test.jpg"))
}

func TestIsDir(t *testing.T) {
	assert.False(t, fileutil.IsDir(""))
	assert.False(t, fileutil.IsDir("/not-exist"))
	assert.True(t, fileutil.IsDir("testdata"))
}

func TestIsAbsPath(t *testing.T) {
	assert.NoError(t, fileutil.DeleteFileIfExist("/not-exist"))
}
