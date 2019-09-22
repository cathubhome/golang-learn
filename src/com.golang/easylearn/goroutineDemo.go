package main

import "fmt"

func main() {

	inputChannel := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			inputChannel <- i
		}
		close(inputChannel)
	}()

	resultChannel := make(chan int, 100)
	workers := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go calucate(inputChannel, workers, resultChannel)
	}

	//等待所有计算的goroutine退出
	go func() {
		for i := 0; i < 10; i++ {
			<-workers
			fmt.Println("wait goroutine:", i, " exited")
		}
		close(resultChannel)
	}()

	for v := range resultChannel {
		fmt.Println(v)
	}

}

func calucate(channel chan int, workers chan bool, resultChannel chan int) {
	for v := range channel {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(v)
			resultChannel <- v
		}
	}
	fmt.Println("exit")
	workers <- true
}
