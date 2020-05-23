package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
	"time"
)

/**
协程与线程：
协程：独立的栈空间，共享堆空间，调度由用户控制；
线程：一个线程可以跑多个协程，协程是轻量级的线程
*/

var (
	resultMap = make(map[int]uint64)
	lock      sync.Mutex
)

func main() {

	setRuntimeCpuNum()

	//goroutine的全局变量与锁机制
	demo()

	//主线程不会等协程执行完，所以这里需要等待，当然这是极其low逼的做法
	time.Sleep(time.Second * 3)

	//此处也要加锁，理论上3s可以执行完所有的协程任务，可是程序并不知晓，terminal使用>go build -race goroutine.go 然后执行goroutine.exe发现Found 2 data race(s)
	//tips:如果使用golang的协程，那么可以加上-race查看是否存在资源竞争
	lock.Lock()
	// To store the keys in slice in sorted order
	var keys []int
	for k, _ := range resultMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the operation you want
	for _, k := range keys {
		fmt.Printf("%d! = %v\n", k, resultMap[k])
	}
	lock.Unlock()
}

/**
设置goroutine运行的cpu核数
*/
func setRuntimeCpuNum() {
	//获取当前操作系统的cpu核数
	cpuMum := runtime.NumCPU()
	//设置goroutine使用的cpu核数，高版本go默认是多核，低版本默认运行的cpu核心为1
	//runtime.GOMAXPROCS(cpuMum)
	fmt.Printf("cpu num:%d\n", cpuMum)

}

func demo() {
	for i := 1; i < 10; i++ {
		go calc(&task{num: i})
	}
}

type task struct {
	num int
}

/**
阶乘计算
*/
func calc(task *task) {
	var result uint64 = 1
	for i := 1; i <= task.num; i++ {
		result *= uint64(i)
	}
	lock.Lock()
	resultMap[task.num] = result
	lock.Unlock()
}
