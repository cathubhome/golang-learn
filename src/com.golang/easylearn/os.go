package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/*
终端读写
操作终端相关文件句柄常量:os.Stdin：标准输入、os.Stdout：标准输出、os.Stderr：标准错误输出
*/

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {

	//打印到终端
	fmt.Fprintf(os.Stdout, "output into console")

	//打印到文件
	openFile, e := os.OpenFile("D:/readme.txt", os.O_CREATE|os.O_WRONLY, 0664)
	if e != nil {
		fmt.Println("open file error", e)
		return
	}
	fmt.Fprint(openFile, "output into file")
	defer openFile.Close()

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	//将字符串格式化输入
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)

	//带缓冲的读写
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}

	//打印命令行参数
	//strings := os.Args
	//for i,v := range strings  {
	//	fmt.Printf("args[%d]=%s\n",i,v)
	//}

	//使用flag作命令行参数解析
	var configPath string
	//指针：将传进来的参数赋值、参数名称、参数默认值、解释说明
	flag.StringVar(&configPath, "c", "", "please input config path:")
	//特别注意：一定要flag.Parse()
	flag.Parse()
	fmt.Printf("configPath:%s\n", configPath)

}
