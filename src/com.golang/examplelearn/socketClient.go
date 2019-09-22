package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

/**
socket客户端

http协议1.1后默认是长连接，如果一个client使用http1.1协议，但又不希望使用长链接，则需要在header中指明connection的值为close
如果server方也不想支持长链接，则在response中也需要明确说明connection的值为close

golang长短链接的处理：https://studygolang.com/articles/6301
*/
func main() {

	//demo1()
	demo2()
}

/**
示例1
*/
func demo1() {

	conn, err := net.Dial("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			return
		}
	}
}

/**
示例2
*/
func demo2() {

	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host: www.baidu.com\r\n"
	msg += "Connection: close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}

}
