package fileutil

import (
	"bytes"
	"os"
)

var (
	// DefaultDirPerm perm and flags for create log file
	DefaultDirPerm   os.FileMode = 0775
	DefaultFilePerm  os.FileMode = 0665
	OnlyReadFilePerm os.FileMode = 0444

	// DefaultFileFlags for create and write
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	// OnlyReadFileFlags open file for read
	OnlyReadFileFlags = os.O_RDONLY
)

// PathExists 给定路径是否存在
func PathExists(path string) bool {
	if path == "" {
		return false
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// IsDir 给定资源路径是否为目录
func IsDir(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return fi.IsDir()
	}
	return false
}

// IsFile 给定资源路径是否为文件
func IsFile(path string) bool {
	if path == "" {
		return false
	}

	if fi, err := os.Stat(path); err == nil {
		return !fi.IsDir()
	}
	return false
}

// IsZipFile 是否为zip文件
func IsZipFile(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
