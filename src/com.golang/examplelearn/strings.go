package main

import (
	"fmt"
	"strconv"
	s "strings"
)

var p = fmt.Println

/**
字符串函数操作：使用strings包
字符串的拼接这种常见情景自然少不了，附带字符串拼接的几种方法的性能分析结果： http://herman.asia/efficient-string-concatenation-in-go
*/
func main() {
	//判断是否含有指定字符串
	p("Contains:  ", s.Contains("test", "es"))
	//判断字符串的出现次数
	p("Count:     ", s.Count("test", "t"))
	//判断是否含有指定的字符串前缀
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	//判断是否含有指定的字符串后缀
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	//查询指定字符串的首次出现位置,没有则返回-1
	p("Index:     ", s.Index("test", "e"))
	//查询指定字符串最后出现的位置,没有则返回-1
	p("LastIndex", s.LastIndex("test", "t"))
	//将切片转为以指定的连接符连接的字符串
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	//将字符串重复指定的次数
	p("Repeat:    ", s.Repeat("a", 5))
	//返回将s中前n个old子串都替换为new的新字符串，如果n<0会替换所有old子串
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	//字符串分割返回切片
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	//转小写
	p("ToLower:   ", s.ToLower("TEST"))
	//转大写
	p("ToUpper:   ", s.ToUpper("test"))
	//去掉字符串首尾空白字符
	p("TrimSpace:", s.TrimSpace(" test  "))
	//去掉字符串首尾指定cut字符
	p("Trim:", s.Trim("test", "t"))
	//返回str空格分隔的所有子串的slice
	p("Field:", s.Fields(" i love you forever"))
	//将整数转为字符串,Itoa:Integer to ascii
	p("strconv.Itoa:", strconv.Itoa(1))
	//将字符串转为整数
	i, e := strconv.Atoi("1")
	if e == nil {
		p("strconv.Atoi", i)
	}
	p()

}
