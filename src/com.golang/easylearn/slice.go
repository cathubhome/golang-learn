package main

import (
	"fmt"
)

/**
切片：可以理解成可变长度的数组,严格的说切片有容量与长度两个属性；切片是数组的一个引用，因此切片是引用类型
*/
func main() {

	//方式1：通过切面字面量申明切片，创建长度与容量都是5的字符串切片
	slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Printf("slice:%v,length:%d,capacity:%d\n", slice, len(slice), cap(slice))

	//方式2：使用make函数，创建一个float64类型的切片。切片的长度和容量相同为5
	x := make([]float64, 5)
	fmt.Printf("x:%v,length:%d,capacity:%d\n", x, len(x), cap(x))
	//同时指定切片长度和容量，特别注意的是如果定义时length>capacity.但是赋值的时候要注意最大的索引仍然是len(capacity)－1，否则程序执行时会报错
	y := make([]float64, 5, 10)
	fmt.Printf("y:%v,length:%d,capacity:%d\n", x, len(y), cap(y))

	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
	}
	fmt.Println(x)

	for i := len(y) - 1; i > 0; i-- { //如果将len()替换为cap()，则会报错panic: runtime error: index out of range
		y[len(y)-1-i] = float64(i)
	}
	fmt.Println(y)

	//创建nil切片
	var nilSlice []int
	fmt.Printf("nil整型切片：%v,nilSlice为nil:%t\n", nilSlice, nilSlice == nil)

	//创建空切片(使用make创建空的整型切片、使用切片字面量创建空的整型切片)
	//空切片在底层数组包含 0 个元素，也没有分配任何存储空间。
	// 想表示空集合时空切片很有用，例如，数据库查询返回 0 个查询结果时
	emptySlice1 := make([]int, 0)
	emptySlice2 := []int{}
	fmt.Printf("emptySlice1:%v\n", emptySlice1)
	fmt.Printf("emptySlice2:%v\n", emptySlice2)

	fmt.Println("通过数组切片赋值")
	//用[low_index:high_index]的方式获取数值切片，其中切片元素包括low_index的元素，但是不包括high_index的元素
	fmt.Println("x[1:3] = ", x[1:3])
	fmt.Println("x[:3] = ", x[:3])
	fmt.Println("x[2:] = ", x[2:])
	fmt.Println("x[:] = ", x[:])

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

	fmt.Println("++++++有趣的：两个切片共享同一个底层数组，如果一个切片修改了该底层数组的共享部分，另一个切片也能感知到+++++++++")

	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice3 := []int{10, 20, 30, 40, 50}
	// 创建一个新切片， 其长度是 2 个元素，容量是 4 个元素
	newSlice := slice3[1:3]
	// 修改 newSlice 索引为 1 的元素,同时也修改了原来的 slice 的索引为 2 的元素
	newSlice[1] = 3
	fmt.Println(slice3)

}
