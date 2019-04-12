package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
随机数
*/
func main() {

	rd := rand.New(rand.NewSource(time.Now().UnixNano())) //为了随机数生成器具有随机性，给一个以当前时间的nano值的动态种子
	intn := rd.Intn(100)
	fmt.Println(intn)
	fmt.Println()

	//0.0 <= f < 1.0
	f := rd.Float64()
	fmt.Println(f)
	fmt.Println()

}
