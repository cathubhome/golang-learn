package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

/**
异常处理：defer、panic与recover
defer来在函数运行结束的时候运行一段代码或调用一个清理函数，最常见的用途在于释放各种资源
*/
func main() {
	file := readFile("D:\\readme.txt")
	fmt.Println(file)
	handleException()
}

//读取文件并返回文件内容
func readFile(filePath string) string {
	f, err := os.Open(filePath)
	//在使用os包中的Open方法打开文件后，立马跟着一个defer语句用来关闭文件句柄，这是编程的良好习惯
	//这样就保证了该文件句柄在main函数运行结束的时候或者异常终止的时候一定能够被释放
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}

	bReader := bufio.NewReader(f)
	buf := bytes.Buffer{}
	for {
		line, ok := bReader.ReadString('\n')
		if ok != nil {
			break
		}
		//去除换行符与字符串拼接，引申：字符串拼接的处理方法-https://blog.csdn.net/guizaijianchic/article/details/78581545
		buf.WriteString(strings.Trim(line, "\r\n"))
	}
	return buf.String()
}

//panic&&recover的使用:panic和recover是Go语言提供的用以处理异常的关键字
//panic用来触发异常，而recover用来终止异常并且返回传递给panic的值。（特别注意recover并不能处理异常，而且recover只能在defer里面使用，否则无效。）
func handleException() {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	fmt.Println("I am walking and singing...")
	panic("It starts to rain cats and dogs")
}
