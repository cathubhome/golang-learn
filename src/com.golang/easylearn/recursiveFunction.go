package main

import "fmt"

/**
递归函数
*/
func main() {
	n := 6
	i := fibonacci(n)
	fmt.Println("斐波拉切数列n为:", n, "计算结果为:", i)
}

//斐波拉切数列:第一个元素是1，第二个元素是2，后面的元素依次是前两个元素之和,{1,2,3,5,8,13...}
/**
  f(1)=1
  f(2)=2
  f(n)=f(n-2)+f(n-1)
*/
func fibonacci(n int) int {
	var retVal = 0
	if n == 1 {
		retVal = 1
	} else if n == 2 {
		retVal = 2
	} else {
		retVal = fibonacci(n-2) + fibonacci(n-1)
	}
	return retVal
}
