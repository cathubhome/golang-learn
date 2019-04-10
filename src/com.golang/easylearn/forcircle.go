package main

/**
for循环
*/
func main() {

	/**
	for直到型循环
	*/
	for i := 1; i <= 100; i++ {
		println(i)
	}

	println("------------------------------")

	/**
	go没有提供while关键字，模拟当型循环如下
	*/
	j := 1
	for j <= 100 {
		println(j)
		j++
	}

}
