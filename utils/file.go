//@Time : 2020/10/14 下午4:25
//@Author : bishisimo
package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func IsExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//递归读取所有子文件
func GetAllFilePath(dir string) ([]string, error) {
	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.New("unable to read path of " + dir)
	}
	// 获取文件，并输出它们的名字
	filePaths := make([]string, 0)
	for _, file := range files {
		filePath := path.Join(dir, file.Name())
		state, err := os.Stat(filePath)
		if err != nil {
			return nil, errors.New("unable to read path of " + filePath)
		} else if state.IsDir() {
			f, err := GetAllFilePath(filePath)
			if err != nil {
				return nil, err
			} else {
				filePaths = append(filePaths, f...)
			}
		} else {
			filePaths = append(filePaths, filePath)
		}
	}
	return filePaths, nil
}

//只读取一层子文件
func GetAllFileSubPath(key string, dir string) (map[string]string, error) {
	filePaths, err := GetAllFilePath(dir)
	if err != nil {
		return nil, err
	}
	splitString := dir
	if splitString[len(splitString)-1:] != "/" {
		splitString += "/"
	}
	fileSubPaths := make(map[string]string)
	for _, filePath := range filePaths {
		subPath := strings.Split(filePath, splitString)
		//fileSubPaths [filePath]=path.Join(key,subPath[len(subPath)-1])
		fileSubPaths[filePath] = subPath[len(subPath)-1]
	}
	return fileSubPaths, nil
}

func MakeDir(dir string) error {
	if !IsExistPath(dir) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
