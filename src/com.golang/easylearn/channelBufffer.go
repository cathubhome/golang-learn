package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
通道缓冲区:
在定义chan变量的时候，还可以指定它的缓冲区大小
一般我们定义的channel都是同步的，也就是说接受端和发送端彼此等待对方ok才开始
但是如果你给一个channel指定了一个缓冲区，那么消息的发送和接受式异步的，除非channel缓冲区已经满了，也就是说缓冲区满了才阻塞程序
*/
func main() {

	var c = make(chan string, 20)
	go shooting(c)
	go count(c)
	var input string
	fmt.Scanln(&input)

}

func shooting(msg_chan chan string) {
	var group = 1
	for {
		for i := 1; i <= 10; i++ {
			//将int转string,strconv.iota()函数
			msg_chan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
		}
		group++
		time.Sleep(time.Second * 3)
	}
}

func count(msg_chan chan string) {
	for {
		fmt.Println(<-msg_chan)
	}
}
