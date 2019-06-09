package main

import (
	"fmt"
	"time"
)

/**
时间与日期
*/
func main() {

	//当前时间戳
	unix := time.Now().Unix()
	fmt.Println("current time timestamp:", unix)

	//当前时间
	now := time.Now()
	fmt.Println("current time:", now)

	//time格式化,"2006-01-02 15:04:05"这个是固定写法,是go诞生之日
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println("format time:", format)

	//时间戳转格式化时间字符串
	s := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
	fmt.Println("timestamp convert to standard time:", s)

}
