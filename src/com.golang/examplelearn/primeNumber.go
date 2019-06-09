package main

import "fmt"

//输出区间的素数,1不是素数
func main() {
	min := 2
	max := 100
	//初级的写法
	numberSlice := getPrimeNumber(min, max)
	fmt.Println(numberSlice)

	//高级的写法

	// The prime sieve: Daisy-chain filter processes together.
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print("prime:", prime, " \n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		fmt.Println("ch = ch1")
	}

}

func getPrimeNumber(min int, max int) []int {
	primeNumberArray := make([]int, 0) //空切片
	for i := min; i <= max; i++ {
		primeNumber := judgePrimeNum(i)
		if primeNumber {
			primeNumberArray = append(primeNumberArray, i)
		}
	}
	return primeNumberArray
}

func judgePrimeNum(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; i <= 100; i++ {
		ch <- i // Send 'i' to channel 'ch'.
		fmt.Println("generate func,ch:", i)
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		fmt.Println("i:", i, "prime:", prime)
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'
			fmt.Println("out <-i", i)
		}
	}
}
