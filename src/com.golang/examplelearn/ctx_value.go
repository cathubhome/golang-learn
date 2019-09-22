package main

import (
	"context"
	"fmt"
)

/*
golang包中context可以用于
①goroutine的超时控制
②保存上下文数据
*/

func processVal(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		fmt.Println("trace_id can not empty")
		return
	}

	fmt.Printf("ret:%d\n", ret)

	s, ok := ctx.Value("session_id").(string)
	if !ok {
		fmt.Printf("session_id can not empty")
		return
	}
	fmt.Printf("session_id:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 13483434)
	ctx = context.WithValue(ctx, "session_id", "AWSEDRFTGYHUJUIKKPM")
	processVal(ctx)
}
