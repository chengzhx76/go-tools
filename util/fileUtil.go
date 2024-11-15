package util

import (
	. "github.com/chengzhx76/go-tools/consts"
	"os"
	"path/filepath"
	"strings"
)

// 判断文件或文件夹是否存在
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
func CreateDirectory(fileDir string) error {
	ext := filepath.Ext(fileDir)
	dir := fileDir
	if !IsBlank(ext) {
		dir = filepath.Dir(dir)
	}
	if !IsExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 创建文件 没有则创建
func CreateFile(pathName string) error {
	dir, _ := GetDirAndFileName(pathName)
	if !IsExist(dir) {
		err := CreateDirectory(dir)
		return err
	}
	file, err := os.Open(pathName)
	defer func() { file.Close() }()
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(pathName)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDirAndFileName(filePath string) (string, string) {
	index := strings.LastIndex(filePath, SYMBOL_SLASH)
	fileName := filePath[index+1:]
	dir := filePath[:index+1]
	return dir, fileName
}
