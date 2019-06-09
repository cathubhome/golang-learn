package main

import "fmt"

/**
结构体：Go提供的结构体就是把使用各种数据类型定义的不同变量组合起来的高级数据类型
		Go使用结构体组合函数的方式来为结构体定义结构体方法
*/

//type关键字（表示要定义一个新的数据类型)+ 结构体名称 + struct关键字
type Rect struct {
	//矩形长与宽
	length, width float64
}

type ColorRect struct {
	Rect  //匿名字段（嵌入字段）,默认colorRect就拥有了rect的所有字段
	color string
}

//计算矩形面积
func calculateRectArea(rect Rect) float64 {
	return rect.length * rect.width
}

//Go函数的参数传递方式是值传递，这句话对结构体同样适用
func deliver(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.length * rect.width
}

//重点：结构体组合函数（结构体“内部函数”）
//首先是关键字func表示这是一个函数，第二个参数是结构体类型和实例变量，第三个是函数名称，第四个是函数返回值
//area()函数和普通函数定义的区别就在于area()函数多了一个结构体类型限定，这样一来Go就知道了这是一个为结构体定义的方法
func (rect Rect) area() float64 {
	return rect.width * rect.length
}

//结构体与指针,使不使用结构体指针和使不使用指针的出发点是一样的，那就是你是否试图在函数内部改变传递进来的参数的值
func (rect *Rect) areaPtr() float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func main() {
	//按照结构体成员定义的顺序来赋值
	rect := Rect{5, 3.0}
	//如下使用 key:value的方式赋值
	//rect := Rect{length:5,width:3.0}
	//结构体类型可以通过.来访问内部的成员，包括给内部成员赋值和读取内部成员值
	//var rect Rect
	//rect.length = 2.5
	//rect.width = 3.0
	fmt.Println("rect area:", calculateRectArea(rect))
	fmt.Println()

	fmt.Println("multi rec area:", deliver(rect))
	fmt.Println("rect length:", rect.length)
	fmt.Println("rect width:", rect.width)
	fmt.Println()

	fmt.Println("After execute area(),calculate result:", rect.area())
	fmt.Println()

	rectangle := new(Rect)
	rectangle.width = 10
	rectangle.length = 10
	fmt.Println("origin address:", *rectangle)
	fmt.Println("Width:", rectangle.width, "Length:", rectangle.length,
		"Area:", rectangle.areaPtr())
	fmt.Println("final address:", *rectangle)
	fmt.Println()

	colorRect := ColorRect{Rect: Rect{width: 1, length: 1}, color: "red"}
	fmt.Println("length:", colorRect.length, "width:", colorRect.length, "color:", colorRect.color)

}
