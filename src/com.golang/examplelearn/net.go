package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
Go提供了一个完善的net/http包，通过http包可以快速搭建起来一个可以运行的Web服务，
同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作
*/
func main() {
	//设置访问的路由
	http.HandleFunc("/", index)
	//设置监听的端口与处理请求和生成返回信息的处理逻辑
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//解析参数，默认是不会解析的
	r.ParseForm()
	//输出到服务器端的打印信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//输出到客户端信息
	fmt.Fprintf(w, "welcome to paic!")
}
