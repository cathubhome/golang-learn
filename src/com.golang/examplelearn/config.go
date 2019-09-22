package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

/**

beego config模块的使用

E:\amusement\learning\goProject>go run src\com.golang\examplelearn\config.go
port: 8080
logLevel: [debug]
logPath: C:\\Windows\\WindowsUpdate.log

*/
func main() {
	conf, err := config.NewConfig("ini", "./config/log.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}

	fmt.Println("port:", port)
	logLevel := conf.Strings("log::logLevel")
	if err != nil {
		fmt.Println("read logLevel failed, ", err)
		return
	}
	fmt.Println("logLevel:", logLevel)

	log_path := conf.String("log::logPath")
	fmt.Println("logPath:", log_path)
}
