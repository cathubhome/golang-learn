package main

import (
	"fmt"
	"os"
	"time"
)

/**
type error interface {
	Error() string
}
*/

/**
自定义错误需实现error接口
*/
type filePathError struct {
	filePath string
	message  string
	time     string
}

/**
实现Error() string 方法
*/
func (filePathError *filePathError) Error() string {
	return fmt.Sprintf("filePath=%s message=%s time=%s", filePathError.filePath, filePathError.message, filePathError.time)
}

func main() {
	error := readFilePath("d:\readme.md")
	if error != nil {
		fmt.Println(error)
	}
}

/**
读取文件
*/
func readFilePath(filePath string) error {
	file, e := os.Open(filePath)
	if e != nil {
		//return errors.New("filePath not found error")
		//return fmt.Errorf("%v",e)
		return &filePathError{
			filePath: filePath,
			message:  e.Error(),
			time:     time.Now().String(),
		}
	}
	defer file.Close()
	return nil
}
