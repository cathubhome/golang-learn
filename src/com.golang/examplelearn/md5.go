package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/**
md5加密
*/
func main() {

	fmt.Println(MD5("1Q2W3E4R5T"))

}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
