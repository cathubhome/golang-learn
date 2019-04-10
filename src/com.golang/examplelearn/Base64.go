package main

import (
	"encoding/base64"
	"fmt"
)

/**
go提供了对base64编码与解码的内置支持
*/
func main() {

	data := "abc123!?$*&()'-=@~"

	//go支持标准与兼容的URL的base64编码，函数入参是byte[],区别在于标准的编码后面是+，而兼容URL的编码方式后面是-
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	//解码一个base64编码可能返回一个错误，
	// 如果你不知道输入是否是正确的base64 编码,你需要检测一些解码错误
	sDec, error := base64.StdEncoding.DecodeString(sEnc)
	if error != nil {
		panic(error)
	} else {
		fmt.Println(string(sDec))
	}
	fmt.Println()

	// 使用兼容URL的base64编码和解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
