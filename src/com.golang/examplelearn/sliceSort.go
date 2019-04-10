package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

//会员（会员）
type Member struct {
	//会员注册名
	Name string `json:"name"`
	//会员积分
	Integral uint8 "json:integral"
	//会员注册时间
	RegisterTime string "json:registerTime"
}

/**
切片排序：Go的sort包实现了内置数据类型和用户自定义数据类型的排序功能
*/
func main() {

	// 这些排序方法都是针对内置数据类型的
	// 这里的排序方法都是就地排序，也就是说排序改变了切片内容，而不是返回一个新的切片
	sorts := []string{"c", "a", "b"}
	sort.Strings(sorts)

	fmt.Println("strings:", sorts)
	// 对于整型的排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("order by aes:", ints)
	// 检测int类型切片是否已经排序好
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("order by desc:", ints)

	//自定义排序，时间排序
	members1 := memberList{
		Member{"Tom", 80, "2019-02-01 12:00:00"},
		Member{"Smith", 60, "2019-01-01 13:00:00"},
		Member{"Tonny", 100, "2019-04-08 22:25:30"}}
	members2 := []Member{
		Member{"Tom", 80, "2019-02-01 12:00:00"},
		Member{"Smith", 60, "2019-01-01 13:00:00"},
		Member{"Tonny", 100, "2019-04-08 22:25:30"},
	}
	fmt.Println(reflect.TypeOf(members1))
	sort.Sort(members1)
	fmt.Println(members1)
	sort.Sort(memberList(members2))
	fmt.Println(members2)

}

type memberList []Member

func (I memberList) Len() int {
	return len(I)
}

func (I memberList) Less(i, j int) bool {
	return stringConvertTime(I[i].RegisterTime) > stringConvertTime(I[j].RegisterTime)
}

func (I memberList) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

func stringConvertTime(timeStr string) int64 {
	//获取时区,"Local"为本地时区
	loc, _ := time.LoadLocation("Local")
	// 转化需要的模板，下面这行代码是固定写法，只能用这个时间字符串,够坑爹的！！！
	timeLayout := "2006-01-02 15:04:05"
	certainTime, _ := time.ParseInLocation(timeLayout, timeStr, loc)
	unix := certainTime.Unix()
	return unix

}
