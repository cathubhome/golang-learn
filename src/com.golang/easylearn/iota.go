package main

import (
	"fmt"
)

/**
Iota是golang的常量计数器，只能在常量表达式中使用；
Iota在const出现时重置为0,在const内部每新增一行常量声明将使得iota计数一次（自增1）（iota可以理解const语句块中的行索引）；
Iota经常用于枚举,如time包对星期与月份的定义；
使用下划线跳值；
*/
const num1 = iota
const num2 = iota + 1
const (
	num3 = iota + 2
	num4
	num5
	_
	_
	num6 = iota
)

//多个常量声明在一行的情况下
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry = iota
)

func main() {

	//编译错误，undefined iota,iota只能在常量的表达式中使用
	//fmt.Println(iota)

	//iota在const出现时初始化为0
	fmt.Println(num1)
	fmt.Println(num2)
	//iota在const内部首行初始化为0
	fmt.Println(num3)
	//沿用生成规则，相当于num4=iota+2,而const内部每新增一行声明iota自增1，所以iota此时为1，num4=3
	fmt.Println(num4)
	//沿用生成规则，相当于num5=iota+2,而const内部每新增一行声明iota自增1，所以iota此时为2，num4=4
	fmt.Println(num5)
	//const内部每新增一行声明iota自增1，所以iota此时为5（iota在const语句块中理解为行索引）
	fmt.Println(num6)

	fmt.Println("——————————")

	fmt.Println(Apple)
	fmt.Println(Banana)
	fmt.Println(Cherimoya)
	fmt.Println(Durian)
	fmt.Println(Elderberry)

}
