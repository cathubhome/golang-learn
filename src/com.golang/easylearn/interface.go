package main

import (
	"fmt"
)

/**
接口:type关键字 + 接口名称 + interface关键字,接口不能包含任何变量
1、没有java中类似implement的关键字，golang中的接口不需要显示的实现，只需要一个变量含有接口的所有方法，那么这个变量
就实现了这个接口（如果一个变量含有多个interface类型的方法，那么这个变量就实现了多个接口；如果一个变量只含有interface的部
分方法，那么这个变量没有实现这个接口）
2、空接口：interface{}没有任何方法，所以任何类型都实现了空接口，也就是说任何类型都可以赋值给空接口
3、一个事物的多种形态，都可以按照统一的接口进行操作（多态）
4、接口可以嵌套：一个接口可以嵌套在另外的接口
5、interface类型默认是一个指针
example:
type ReadWrite interface {
               Read(b Buffer) bool
               Write(b Buffer) bool
}
type Lock interface {
               Lock()
               Unlock()
}
type File interface {
               ReadWrite
               Lock
               Close()
}
*/
type Phone interface {
	call()

	sale() int
}

type NokiaPhone struct {
	price int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia,I can call you")
}

func (nokiaPhone NokiaPhone) sale() int {
	return nokiaPhone.price
}

type IPhone struct {
	price int
}

func (iphone IPhone) call() {
	fmt.Println("I am iphone,I can call you")
}

func (iphone IPhone) sale() int {
	return iphone.price
}

//接口类型可以作为结构体的数据成员
type Person struct {

	//手机切片
	phones []Phone

	//姓名
	name string
}

func main() {

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	//判断变量是否实现了指定的接口
	if i, ok := phone.(Phone); ok {
		fmt.Println("implement phone", i.sale())
	}

	phone = new(IPhone)
	phone.call()

	//phones := []Phone{
	//	NokiaPhone{6000},
	//	IPhone{8000},
	//}
	var phones []Phone
	phones = append(phones, NokiaPhone{6000}, IPhone{8000})
	var totalSales = 0
	for _, elem := range phones {
		totalSales += elem.sale()
	}
	fmt.Println("totalPrice:", totalSales)

	person := Person{phones, "Tom"}
	fmt.Println(person)

}
