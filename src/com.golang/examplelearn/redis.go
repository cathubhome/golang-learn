package main

/**
redis作为高性能的nosql数据库，单机即可支持15万qps,通常适合作为缓存使用
使用第三方开源库操作redis:github.com/garyburd/redigo/redis
*/
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "192.168.17.18:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	//auth after connecting redis if you set password
	err = c.Send("auth", "cathome")
	if err != nil {
		fmt.Println("redis auth error:", err)
		return
	}

	//字符串
	//if reply equals "OK",that means set key succeed
	_, err = c.Do("SET", "programming_language", "golang")
	if err != nil {
		fmt.Println("set string failure,error:", err)
	}

	//if you wanna get string value,you should use redis.Sting() function to convert
	s, err := redis.String(c.Do("GET", "programming_language"))
	if err != nil {
		fmt.Println("get key error:", err)
		return
	}
	fmt.Println("key:", "programming_language,", " value:", s)

	//哈希hash表
	_, err = c.Do("HSET", "employee", "age", 20, "job", "programmer", "sex", "male")
	if err != nil {
		fmt.Println("set hset error:", err)
		return
	}
	i, err := redis.Int(c.Do("HGET", "employee", "age"))
	if err != nil {
		fmt.Println("redis get hash error:", err)
		return
	}
	fmt.Println("hash key:employee,age:", i)

	//批量set
	_, err = c.Do("MSet", "book", "golang实战", "website", "https://studygolang.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Strings(c.Do("MGet", "book", "website"))
	if err != nil {
		fmt.Println("HGet failed,", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}

	//列表list（队列操作）
	_, err = c.Do("lpush", "book_list", "golang实战", "golang并发之道", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
	i2, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(i2)

	//expire time过期时间,unit:秒
	_, err = c.Do("expire", "book", 10)
	if err != nil {
		fmt.Println("expire error:", err)
		return
	}

	defer c.Close()
}
