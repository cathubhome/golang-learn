package main

import (
	"fmt"
)

/**
函数:func关键字+函数名+参数列表+返回值列表(注意：在只返回一个返回值时返回值列表两边的括号可以省略)
	 go不支持函数重载，一个包中不能存在有两个重名的函数
	 go支持函数返回多个值，不必如java去定义对象或结构体
	 go支持可变长参数列表（注意：可变长参数定义只能是函数的最后一个参数，与java规定一致!）
     go的函数也是一种类型,可以赋值给变量

在函数内部声明的变量是局部变量，生命周期仅限于函数内部；在函数外部声明的变量是全局变量，声明周期作用于整个包，如果是大写
则作用于整个程序

函数参数的传递方式：
1、值传递
2、引用传递
notice one:
	无论是值传递还是引用传递，传递的都是变量的副本，不过值传递是值的拷贝、引用传递是地址的拷贝，一般来说，地址的拷贝
更为高效，而值的拷贝取决于拷贝的对象的大小，对象越大，性能越低;
notice two:
	map、slice、chan、指针、interface默认以引用的方式传递
*/

//对数组求和
func slice_sum(arr []int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}
	return sum
}

//有趣的是你甚至可以为返回值预先定义一个名称，在函数结束的时候，直接一个return就可以返回所有的预定义返回值，如此处的sum
func slice_sum_another(arr []int) (sum int) {
	sum = 0
	for _, element := range arr {
		sum += element
	}
	return
}

//返回数组之和与平均值
func slice_sum_avg(arr []int) (int, float64) {
	sum := 0
	avg := 0.0
	for _, elem := range arr {
		sum += elem
	}
	avg = float64(sum) / float64(len(arr))
	return sum, avg
}

//求基数与可变长度
func change_len_param_sum(base int, arr ...int) int {
	sum := base
	for _, val := range arr {
		sum += val
	}
	return sum
}

func slice_sum_avg_another(arr []int) (sum int, avg float64) {
	sum = 0
	avg = 0.0
	for _, elem := range arr {
		sum += elem
	}
	avg = float64(sum) / float64(len(arr))
	return sum, avg
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := slice_sum(array)
	fmt.Println("sum:", sum)

	fmt.Println()
	sumAnother := slice_sum_another(array)
	fmt.Println("sum:", sumAnother)

	fmt.Println()
	i, f := slice_sum_avg(array)
	fmt.Println("sum:", i, "avg:", f)

	fmt.Println()
	sum2, avg := slice_sum_avg_another(array)
	fmt.Println("sum:", sum2, "avg:", avg)

	fmt.Println()
	paramSum := change_len_param_sum(10, array...) //使用...将切片打散
	fmt.Println("sum:", paramSum)

	fmt.Println("闭包函数的使用")
	var base = 0
	//闭包函数：所谓闭包就是就是将整个函数一气呵成写好并赋值给一个变量，然后用整个变量作为函数名去调用函数体，其实闭包函数也没什么特别之处
	//由于go不支持在一个函数的内部再定义一个嵌套函数，所以使用闭包能够实现在一个函数内部定义另一个函数的目的
	//特别注意的是闭包函数对它的外层函数具有访问与修改的权限
	total := func(arr ...int) int {
		sum := 0
		for _, val := range arr {
			sum += val
			base++
		}
		return sum
	}
	fmt.Println(total(array...))
	fmt.Println("base:", base)

	fmt.Println("有趣的闭包函数,值得思考哦")
	generator := createEvenGenerator()
	fmt.Println(generator())
	fmt.Println(generator())

}

//有趣的闭包函数，依次返回偶数
//如下定义了一个返回函数定义的函数,而所返回的函数定义就是在这个函数的内部定义的闭包函数,其中func() uint就是函数createEvenGenerator的返回值
func createEvenGenerator() func() uint {
	i := uint(0)
	return func() (retVal uint) {
		retVal = i
		i += 2
		return
	}
}
