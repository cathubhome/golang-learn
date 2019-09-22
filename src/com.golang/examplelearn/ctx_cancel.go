package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
golang包中context可以用于
①goroutine的超时控制
②保存上下文数据
*/

type Result struct {
	r   *http.Response
	err error
}

func processDeal() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("Timeout! err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}
func main() {
	processDeal()
}