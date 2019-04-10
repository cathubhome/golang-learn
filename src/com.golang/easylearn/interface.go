package main

import (
	"fmt"
)

/**
接口:type关键字 + 接口名称 + interface关键字
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
