package fileutil

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"path/filepath"
)

// OpenFile 打开文件，如果不存在则创建
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)

	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// CreateFile 如果不存在，则创建
func CreateFile(fpath string, filePerm, dirPerm os.FileMode) (*os.File, error) {
	dirPath := path.Dir(fpath)
	if !IsDir(dirPath) {
		err := os.MkdirAll(dirPath, dirPerm)
		if err != nil {
			return nil, err
		}
	}

	return os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, filePerm)
}

// WriteFile 写入字符串到文件中
func WriteFile(filePath string, contents string) (int, error) {
	// create and open file
	dstFile, err := OpenFile(filePath, DefaultFileFlags, DefaultFilePerm)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	return dstFile.WriteString(contents)
}

// CopyFile 拷贝文件到另一个文件路径
func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := OpenFile(dstPath, DefaultFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// DeleteIfExist 移除文件或者空目录
func DeleteIfExist(fPath string) error {
	if PathExists(fPath) {
		return os.Remove(fPath)
	}
	return nil
}

// DeleteFileIfExist 如果存在文件，则删除
func DeleteFileIfExist(fPath string) error {
	if IsFile(fPath) {
		return os.Remove(fPath)
	}
	return nil
}

// Unzip zip文件解压缩
func Unzip(archive, targetDir string) (err error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(targetDir, DefaultDirPerm); err != nil {
		return
	}

	for _, file := range reader.File {
		fullPath := filepath.Join(targetDir, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(fullPath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			_ = fileReader.Close()
			return err
		}

		_, err = io.Copy(targetFile, fileReader)

		// close all
		_ = fileReader.Close()
		targetFile.Close()

		if err != nil {
			return err
		}
	}

	return
}
