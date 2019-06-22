package main

import (
	"fmt"
	"reflect"
)

/**
反射：可以在运行时动态的获取变量的相关信息
*/
func main() {

	var x float64 = 3.4
	//reflect.TypeOf，获取变量的类型，返回reflect.Type类型
	fmt.Println("type:", reflect.TypeOf(x))
	//reflect.ValueOf，获取变量的值，返回reflect.Value类型
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	//获取变量的类别，返回一个常量
	fmt.Println("kind:", v.Kind())
	//获取变量的值
	fmt.Println("value:", v.Float())
	//转换成interface{}类型
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	//类型断言
	y := v.Interface().(float64)
	fmt.Println(y)

	//v.SetFloat(6.6) //panic: reflect: reflect.Value.SetFloat using unaddressable value,崩溃了

	//解决方法：传地址，v.Elem()用来获取指针指向的变量，相当于var a *int ; *a = 100
	reflect.ValueOf(&x).Elem().SetFloat(6.6)
	fmt.Println("reflect.ValueOf(x).Elem()获取指针指向的变量", x)

	//用反射操作结构体
	// reflect.Value.NumField()获取结构体中字段
	// reflect.Value.Method(n).Call来调用结构体中的方法

	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	fmt.Println(typ)

	knd := value.Kind() // struct
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

}

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
