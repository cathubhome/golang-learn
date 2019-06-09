package main

import (
	"fmt"
)

/**
切片：可以理解成可变长度的数组,严格的说切片有容量与长度两个属性
*/
func main() {

	//方式1：只指定长度，这个时候切片的长度和容量是相同的
	x := make([]float64, 5)
	//方式2：同时指定切片长度和容量，特别注意的是如果定义时length>capacity.但是赋值的时候要注意最大的索引仍然是len(y)－1，否则程序执行时会报错
	y := make([]float64, 5, 10)
	fmt.Println("Capcity:", cap(x), "Length:", len(x))
	fmt.Println("Capcity:", cap(y), "Length:", len(y))

	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
	}
	fmt.Println(x)

	for i := len(y) - 1; i > 0; i-- { //如果将len()替换为cap()，则会报错panic: runtime error: index out of range
		y[len(y)-1-i] = float64(i)
	}
	fmt.Println(y)

	fmt.Println("通过数组切片赋值")
	//用[low_index:high_index]的方式获取数值切片，其中切片元素包括low_index的元素，但是不包括high_index的元素
	fmt.Println("x[1:3]:", x[1:3])
	fmt.Println("x[:3]:", x[:3])
	fmt.Println("x[2:]:", x[2:])
	fmt.Println("x[:]:", x[:])

	fmt.Println("execute append()")
	//使用append函数给切片增加元素
	x = append(x, 5, 6)
	//Go在默认的情况下，如果追加的元素超过了容量大小，Go会自动地重新为切片分配容量，容量大小为原来的两倍
	fmt.Println("Capcity:", cap(x), "Length:", len(x))
	fmt.Println(x)

	fmt.Println("execute copy()")
	//copy函数用来从一个切片拷贝元素到另一个切片,特别注意的是这里是将后一个参数的切片拷贝到前一个切片,由于y的长度为5，所以最多拷贝5个元素
	copy(y, x)
	fmt.Println("Capcity:", cap(y), "Length:", len(y))
	fmt.Println(y)
}
