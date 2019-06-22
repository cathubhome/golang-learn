package main

import "fmt"

/**
类型断言：由于接口是一般类型，不知道具体类型，如果要转换成具体的类型可以采用类型断言转换
*/
func main() {

	var t int
	var x interface{}
	x = t
	i := x.(int) //转成int
	fmt.Println(i)

	y, ok := x.(int) //转成int,带类型检查
	if ok {
		fmt.Println(y)
	}

	classifier(t)

	//判断变量是否实现了某个接口可以参考interface.go

}

/**
判断参数类型
*/
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}
