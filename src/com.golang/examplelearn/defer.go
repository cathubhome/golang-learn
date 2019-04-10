package main

import (
	"fmt"
	"os"
)

/**
Defer 用来保证一个函数调用会在程序执行的最后被调用。通常用于资源清理工作
*/

//示例：创建文件，然后写入数据，最后关闭文件
func main() {

	// 在使用createFile得到一个文件对象之后，我们使用defer
	// 来调用关闭文件的方法closeFile，这个方法将在main函数
	// 最后被执行，也就是writeFile完成之后
	folderPath := "E:\\temp"
	_, e := isFileExist(folderPath)
	if e != nil {
		fmt.Println("get dir error", e)
	}
	file := createFile(folderPath + "\\" + "defer.txt")
	defer closeFile(file)
	writeFile(file, "this is how to use defer example")
}

//创建文件
func createFile(filePath string) *os.File {
	fmt.Println("create file...")
	file, e := os.Create(filePath)
	if e != nil {
		panic(e)
	}
	return file
}

//向文件写数据
func writeFile(file *os.File, word string) {
	fmt.Println("write file...")
	fmt.Fprint(file, word)
}

//关闭文件
func closeFile(file *os.File) {
	fmt.Println("close file...")
	file.Close()
}

//判断文件夹是否存在,不存在则创建
func isFileExist(folderPath string) (bool, error) {
	_, err := os.Stat(folderPath)
	//文件夹存在时
	if err == nil {
		return true, nil
	}
	//文件夹不存在时创建文件夹
	if os.IsNotExist(err) {
		mkdirError := os.Mkdir(folderPath, os.ModePerm)
		if mkdirError != nil {
			return false, mkdirError
		}
		return true, nil
	}
	return false, err
}
