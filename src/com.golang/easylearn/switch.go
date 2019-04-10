package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
switch说明如下：
(1) switch的判断条件可以为任何数据类型；
(2) 每个case后面跟的是一个完整且独立的程序块，该程序块不需要{}，也不需要break结尾；
(3) 可以为switch提供一个默认选项default，在上面所有的case都没有满足的情况下，默认执行default后面的语句
*/
func main() {

	//使用time包设置时间种子
	rand.Seed(time.Now().UnixNano())
	//生成随机数使用math/rand包,intn为[0,100]随机数，注：不设置时间种子的话，每次生成的rand值会相同
	intn := rand.Intn(100)
	fmt.Println("rand number is ", intn)
	switch intn / 10 {
	case 10:
		fmt.Println("超神")
	case 9:
		fmt.Println("优秀")
	case 8:
		fmt.Println("良好")
	case 7:
		fmt.Println("一般")
	case 6:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")

	}

}
