package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//在terminal中使用go build net.go 运行exe可执行文件

/**
Go提供了一个完善的net/http包，通过http包可以快速搭建起来一个可以运行的Web服务，
同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作
*/
func main() {
	//设置访问的路由
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	//设置监听的端口与处理请求和生成返回信息的处理逻辑
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	//解析url传递的参数,对于post则解析响应包的主体request body，默认是不会解析的
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

func login(w http.ResponseWriter, r *http.Request) {
	//默认情况下,Handler里面是不会自动解析form,必须显式的调用r.ParseForm()后才能对表单数据进行操作
	r.ParseForm()
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		/**
		  Tips: Request本身也提供了FormValue()函数来获取用户提交的参数。
		  如r.Form["username"]也可写成r.FormValue("username")。
		  调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。
		  r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串
		*/
		username := r.FormValue("username")
		password := r.FormValue("password")
		//必填字段校验，
		/**
		tip:r.Form对不同类型的表单元素的留空有不同的处理，
			对于空文本框、空文本区域以及文件上传，元素的值为空值,
			而如果是未选中的复选框和单选按钮，则根本不会在r.Form中产生相应条目，
			如果我们用上面例子中的方式去获取数据时程序就会报错。
			所以我们需要通过r.Form.Get()来获取值，因为如果字段不存在，通过该方式获取的是空值。
			但是通过r.Form.Get()只能获取单个的值，
			如果是map的值，必须通过上面的方式来获取。
		*/
		if len(username) == 0 {
			fmt.Fprint(w, "username can not empty")
		}
		if len(password) == 0 {
			fmt.Fprint(w, "password can not empty")
		}
		//数字校验，校验是数字而不是乱七八糟的字符串
		verificationCode, err := strconv.Atoi(r.Form.Get("verificationCode"))
		if err != nil {
			fmt.Fprint(w, "verificationCode empty or illegal error")
		}
		fmt.Println("verificationCode:", verificationCode)
	}
}

// 处理upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
