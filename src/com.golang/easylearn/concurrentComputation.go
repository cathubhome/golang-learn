package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
并行计算：涉及到数据类型chan与go关键字
		  go关键字用来开启协程（所谓协程，就是go提供的轻量级的独立运算过程，比线程还轻量，创建协程很简单，就是go关键字加上所要运行的函数）计算；
		  channel通道是go用来协程之间通信的方式以及运行同步的机制
*/
func main() {

	sum := computateSum(10)
	println(sum)

	//没有任何输出,这是由于main函数在创建完协程后就立刻退出了，所以协程还没有来得及运行,
	channelDemo1()
	//要求用户输入任何数据后才退出，这样协程就有了运行的时间
	var input string
	fmt.Scanln(&input)
}

//demo1:给定一个整数N，计算从1到N之间既能被3整除又能被5整除的数之和
func computateSum(n int) (sum int) {
	resultChan := make(chan int, 3)
	start := time.Now()
	go getSumOfDivisible(n, 3, resultChan)
	go getSumOfDivisible(n, 5, resultChan)
	//这里其实哪个是被3整除，哪个是被5整除看具体调度方法，不过由于是求和，所以没关系（这部分并行计算的）
	sum3, sum5 := <-resultChan, <-resultChan
	//单独算被15整除的
	go getSumOfDivisible(n, 15, resultChan)
	sum15 := <-resultChan
	end := time.Now()
	fmt.Println("spend time:", end.Sub(start))
	sum = sum3 + sum5 - sum15
	return
}

//最后一个参数是整型chan类型：可以理解成一个FIFO队列，其所能接收的数据类型是由chan关键字后面的类型所决定的
func getSumOfDivisible(num int, divider int, resultChan chan int) {
	sum := 0
	for value := 0; value < num; value++ {
		if value%divider == 0 {
			sum += value
		}
		fmt.Println("divider:", divider)
		//随机睡眠一段时间,验证并行计算部分的确是并行的
		tick := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * tick)
	}
	// <- 表示将函数的计算结果写入channel,channel是go提供的用来协程之间通信的方式
	resultChan <- sum
}

//channel通道demo: 学员投篮，教练计数
// 特别注意：如果你要向channel里面写信息，必须有配对的取信息的一端，否则是不会写，如果将go count(msgChan)注释掉，那么程序就不再输出消息与提示消息了
func channelDemo1() {
	var msgChan chan string
	msgChan = make(chan string)
	go shooting(msgChan)
	go count(msgChan)
}

func shooting(msgChan chan string) {
	for {
		msgChan <- "投篮" //向channel写入数据
		fmt.Println("继续投篮")
	}
}

func count(msgChan chan string) {
	for {
		msg := <-msgChan //向channel读取数据
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}
