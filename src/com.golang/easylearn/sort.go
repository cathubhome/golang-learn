package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
排序:
sort.Ints对整数进行排序
sort.Strings对字符串进行排序
sort.Float64s对浮点数进行排序

当这个函数在目标slice中搜索不到被搜索元素时,不同于java返回-1，
而是返回了被搜索的元素应该在目标slice中按升序排序时该被搜索元素插入的位置
所以还得判断一下目标slice中返回值这个位置到底是不是正在搜索的这个元素

sort.SearchInts(a []int, b int) 从数组a中查找b，特别注意前提是a必须有序

*/
func main() {

	//随机种子
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice1 := make([]int, 0)
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, rd.Intn(10))
	}
	//判断切片是否有序
	sorted := sort.IntsAreSorted(slice1)
	if !sorted {
		sort.Ints(slice1)
	}
	fmt.Printf("slice1:%v\n", slice1)
	randNum := rd.Intn(10)
	fmt.Printf("random num:%d\n", randNum)
	position := sort.SearchInts(slice1, randNum)
	fmt.Printf("position:%d\n", position)
	if slice1[position] == randNum {
		fmt.Printf("%d exist in slice1 %v \n", randNum, slice1)
	} else {
		fmt.Printf("%d not exist in slice1 %v \n", randNum, slice1)
	}

}

/**
冒泡排序:
*/
func bubbleSort(slice []int) {

	for i := 0; i < len(slice); i++ {
		for j := 1; j < len(slice)-i; j++ {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}

}
