package main

import "fmt"

/**
数组说明：相同数据类型的元素组成的固定长度的有序集合
*/
func main() {

	//第一种数组定义方法，也是最基本的，使用var关键字来定义，然后依次给元素赋值。
	// 对于没有赋值的元素，默认为零值(对于整数，零值就是0，浮点数，零值就是0.0，字符串，零值就是””，对象零值就是nil)
	var x [3]int //显式定义了数组的长度,同var x [3]int{}
	x[0] = 2
	x[1] = 3
	var sum int
	//range函数可以用在数组，切片和字典上面。当range来遍历数组的时候返回数组的索引和元素值,这里是对元素求和，并不关心索引，在go中当你对一个函数的返回值不感兴趣时使用下划线（_）替代
	// 另外如果定义了索引（将下划线替换成变量i），却在循环中没有用到索引，这时go在编译时会报错
	for _, elem := range x {
		sum += elem
	}
	for i, elem := range x {
		sum += elem
		fmt.Println("current index is ", i)
	}
	fmt.Println(sum)

	//第二种方式：使用...代替数组长度，go会自动推断数组长度，不过这种定义方式必须有初始化的值
	//两种定义方法如下
	//var y = [...]string{
	//	"Monday",
	//	"Tuesday",
	//	"Wednesday",
	//	"Thursday",
	//	"Friday",
	//	"Saturday",
	//	"Sunday"}
	var y = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	for _, elem := range y {
		println(elem)
	}

}
