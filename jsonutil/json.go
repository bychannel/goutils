package jsonutil

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// WriteFile 写数据到json文件中
func WriteFile(filePath string, data interface{}) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, jsonBytes, 0664)
}

// ReadFile 读取json文件数据
func ReadFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(v)
}
