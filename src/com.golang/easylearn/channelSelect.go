package main

import (
	"fmt"
	"time"
)

/**
select多通道示例1：如果上面的投篮训练现在有两个教练了，各自负责一个训练项目，而且还在不同的篮球场，这个时候很显然，我们一个channel就不够用了，这是select就派上了用场
*/
func main() {
	c_fixed := make(chan string)
	c_3_point := make(chan string)
	go fixed_shooting(c_fixed)
	go three_point_shooting(c_3_point)
	//我们将定点投篮和三分投篮的消息写入了不同的channel，那么main函数如何知道从哪个channel读取消息呢？
	//使用select方法，select方法依次检查每个channel是否有消息传递过来，如果有就取出来输出。
	//如果同时有多个消息到达，那么select闭上眼睛随机选一个channel来从中读取消息，
	//如果没有一个channel有消息到达，那么select语句就阻塞在这里一直等待
	//go func() {
	//	for {
	//		select {
	//		case msg1 := <-c_fixed:
	//			fmt.Println(msg1)
	//		case msg2 := <-c_3_point:
	//			fmt.Println(msg2)
	//		}
	//	}
	//}()
	//在某些情况下，比如学生投篮中受伤了，那么就轮到医护人员上场了，
	//教练在一旁看看，如果是重伤，教练就不等了，就回去休息了，待会儿再过来看看情况。
	//我们可以给select加上一个case用来判断是否等待各个消息到达超时
	go func() {
		for {
			select {
			case msg1 := <-c_fixed:
				fmt.Println(msg1)
			case msg2 := <-c_3_point:
				fmt.Println(msg2)
				//为什么各个channel消息都没有到达，select为什么不阻塞？
				//就是因为这个time.After，虽然它没有显式地告诉你这是一个channel消息，但是记得么？main函数也是一个channel啊！
				//至于time.After的功能实际上让main阻塞了5秒后返回给main的channel一个时间，所以我们在case里面把这个时间消息读出来，select就不阻塞了
			case <-time.After(time.Second * 5):
				fmt.Println("timeout, check again...")
				//select的default选项，当select发现没有消息达到时也不会阻塞,直接跳转回去再次判断
				//default:
				//	fmt.Println("cannot find any msg!")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

func fixed_shooting(msg_chan chan string) {
	//学员定点投篮三次后受伤了
	var times = 3
	var t = 1
	for {
		if t <= times {
			msg_chan <- "fixed shooting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func three_point_shooting(msg_chan chan string) {
	//学员三分投篮五次后受伤了
	var times = 5
	var t = 1
	for {
		if t <= times {
			msg_chan <- "three point shooting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}
