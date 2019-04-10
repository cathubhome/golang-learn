package main

import (
	"fmt"
)

/**
字典map:一组无序的键值对的集合
*/
func main() {

	//[]之间的是键类型，右边的是值类型
	hashmap := map[string]string{
		"A": "applce",
		"B": "banana",
		"O": "orange",
	}
	for key, val := range hashmap { //range函数用来迭代字典元素，返回key:value键值对
		fmt.Println("key:", key, "value:", val)
	}

	println("make函数初始化字典")

	//如果不用make函数初始化字典，则在程序执行后会报错：panic: assignment to entry in nil map
	x := make(map[string]int64)
	x["A"] = 0
	x["B"] = 1
	x["C"] = 2
	fmt.Println(x)
	//如果key不存在则会返回值类型的零值，这里如何知道key存在且值为零还是不存在呢？
	fmt.Println("x['D']:", x["D"])

	fmt.Println("判断字典的key是否存在")
	if val, ok := x["D"]; ok {
		fmt.Println("key exist,key-value:", val)
	} else {
		fmt.Println("key not exist!")
	}

	fmt.Println("删除键值对")
	delete(x, "A")
	fmt.Println(x)

	fmt.Println("practice练习")
	//场景：学生登记表，登记表中有一组学号，每个学号对应学生，每个学生有名字与年龄
	var facebook = make(map[string]map[string]int)
	facebook["0616020432"] = map[string]int{"Jemy": 25}
	facebook["0616020433"] = map[string]int{"Andy": 23}
	facebook["0616020434"] = map[string]int{"Bill": 22}
	for stu_no, stu_info := range facebook {
		fmt.Println("Student:", stu_no)
		for name, age := range stu_info {
			fmt.Println("Name:", name, "Age:", age)
		}
	}

}
