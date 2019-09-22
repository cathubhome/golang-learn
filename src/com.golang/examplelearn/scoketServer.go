package main

import (
	"fmt"
	"net"
)

/**
tcp服务端:todo 建议将netty与golang的tcp编程对比下，深入研究下长短链接以及拆粘包的问题，golang是如何解决拆粘包问题的也需要深入学习
*/
func main() {
	fmt.Println("start server ...")
	listener, e := net.Listen("tcp", "0.0.0.0:5000")
	if e != nil {
		fmt.Println("listen failure,error:", e)
		return
	}
	for {
		//监听到新的连接，创建goroutine交给process函数处理
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("accept failure,error:", e)
			continue
		}
		go process(conn)
	}

}

/**
处理连接
*/
func process(conn net.Conn) {
	//处理完毕,关闭链接,否则存在句柄泄漏
	defer conn.Close()
	fmt.Println("client address:", conn.RemoteAddr())
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		fmt.Println("read:", string(buf))
	}
}
