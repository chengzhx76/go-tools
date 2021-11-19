package tool

import (
	"os"
)

//判断文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 创建目录 没有则创建
func CreateDirectory(path string) error {
	if !IsExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 创建目录 没有则创建
func CreateFile(pathName string) error {
	file, er := os.Open(pathName)
	defer func() { file.Close() }()
	if er != nil && os.IsNotExist(er) {
		file, _ = os.Create(pathName)
	}
	return nil
}
