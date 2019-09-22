package main

import (
	"encoding/json"
	"fmt"
)

/**
go内置对json的支持:使用两个结构体来演示自定义数据类型的JSON数据编码和解码
序列化：json.Marshal(data interface{})
反序列化：json.Unmarshal(data []byte,v interface{})
*/

type Response1 struct {
	Id       int
	Page     int
	Fruits   []string
	shopping string //成员小写时不可导出
}

type Response2 struct {
	Page   int      `json:"page"` //使用tag自定义json键的名称
	Fruits []string `json:"fruits"`
}

func main() {

	//将切片编码为json数组
	slcD := []string{"apple", "peach", "pear"}
	slcB, error := json.Marshal(&slcD)
	if error != nil {
		fmt.Println("Marshal failure!", "生成json字节数组失败")
	} else {
		//slcB是byte[]类型，转string类型便于查看
		fmt.Println(string(slcB))
	}

	//将字典编码为json对象
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(&mapD)
	fmt.Println(string(mapB))

	//json包可以自动编码自定义数据类型，结果将只包括自定义类型中可导出成员的值，并且在默认情况下，这些成员名称作为json数据的键
	var rest1 Response1
	rest1.Fruits = []string{"apple", "orange"}
	rest1B, _ := json.Marshal(&rest1)
	fmt.Println(string(rest1B))

	//使用tag来自定义编码后的json键的名称
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(&res2D)
	fmt.Println(string(res2B))

	//解码json数据为go数据
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	fmt.Println("byt:", byt)
	// 我们需要提供一个变量来存储解码后的JSON数据，这里
	// 的`map[string]interface{}`将以Key-Value的方式
	// 保存解码后的数据，Value可以为任意数据类型
	var dat map[string]interface{}
	// 解码过程，并检测相关可能存在的错误，特别注意这里的json.Unmarshal(byt, &dat)不能写成json.Unmarshal(byt, &dat)
	// 否则报错：panic: json: Unmarshal(non-pointer map[string]interface {})
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("dat:", dat)

	//为了使用解码后map里面的数据，我们需要将Value转换为它们合适的类型
	num := dat["num"].((float64))
	fmt.Println("value:", num)

	// 访问嵌套的数据需要一些类型转换
	strs := dat["strs"].([]interface{})
	if len(strs) > 0 {
		for i, elem := range strs {
			fmt.Println("index:", i, "value:", elem.(string))
		}
	}

	// 我们还可以将JSON解码为自定义数据类型，这有个好处是可以
	// 为我们的程序增加额外的类型安全并且不用再在访问数据的时候
	// 进行类型断言
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

}
