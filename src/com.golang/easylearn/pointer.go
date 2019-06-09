package main

import (
	"fmt"
)

/**
指针（指针变量）：所谓指针可以理解成一个箭头，这个箭头就是指向（存储）一个变量的地址，由于箭头本身也需要变量来存储，所以也称之为指针变量（注意：go不支持乱七八糟的指针移位,不同C语言支持运算）
*/
func main() {

	demo1()
	fmt.Println()

	fmt.Println("change()")
	x := 100
	change(&x)
	fmt.Println("After change(),x:", x)
	fmt.Println()

	fmt.Println("newInitPtr()")
	newInitPtr()
	fmt.Println()

	fmt.Println("swap()")
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("After swap,a:", a)
	fmt.Println("After swap,b:", b)
	fmt.Println()

}

//指针的使用demo
func demo1() {

	// 变量x
	x := 10
	// 指针变量，指向x变量的地址（&就是取一个变量的地址）
	xPtr := &x //同var xPtr *int = x
	fmt.Println("x:", x)
	fmt.Println("x_ptr:", xPtr)
	// 通过指针变量输出x的值（* 就是取一个指针变量所指地址的值）
	fmt.Println("*x_ptr", *xPtr)

}

//指针的一大用途就是可以将变量的指针作为实参传递给函数，从而在函数内部能够直接修改实参所指向的变量值
func change(x *int) {
	//修改实参所指向的变量值
	*x = 200
	//特别注意：上面的代码*x = 200不同于下面两行代码，下面的两行代码是错误示范，也就是说下面两行代码仅仅是改变内部变量x的值，而不会改变传递进来的实参
	//tmp := 200
	//x = &tmp
}

//使用new函数初始化一个指针
func newInitPtr() {
	xPtr := new(int)
	change(xPtr)
	fmt.Println("xPtr指向的地址", xPtr)
	fmt.Println("xPtr本身的地址", &xPtr)
	fmt.Println("xPtr指向的地址值", *xPtr)
}

//交换两变量的值，交叉赋值
func swap(x, y *int) {
	*x, *y = *y, *x
}
