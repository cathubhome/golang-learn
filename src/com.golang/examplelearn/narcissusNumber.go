package main

import (
	"fmt"
	"math"
)

/**
水仙花数：“水仙花数”是指一个三位数,其各位数字立方和等于该数本身
*/
func main() {

	for i := 100; i < 1000; i++ {
		if isNarcissusNumber(i) {
			fmt.Println(i)
		}
	}
}

/**
判断是否是水仙花数
*/
func isNarcissusNumber(num int) bool {
	i := num / 100     //百位数
	j := num / 10 % 10 //十位数
	z := num % 10      //个位数
	temp := (int)(math.Pow(float64(i), 3) + math.Pow(float64(j), 3) + math.Pow(float64(z), 3))
	if num == temp {
		return true
	}
	return false
}
