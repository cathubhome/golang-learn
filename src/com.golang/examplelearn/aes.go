package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/**
aes加密
*/
func main() {
	message := []byte("Hello!My name is X.")
	//指定密钥
	key := []byte("14725836qazwsxed")
	//加密
	cipherText := AEC_CRT_Crypt(message, key)
	fmt.Println("加密后为：", string(cipherText))
	//解密
	plainText := AEC_CRT_Crypt(cipherText, key)
	fmt.Println("解密后为：", string(plainText))
}

func AEC_CRT_Crypt(text []byte, key []byte) []byte {
	//指定加密、解密算法为AES，返回一个AES的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//指定计数器,长度必须等于block的块尺寸
	count := []byte("12345678abcdefgh")
	//指定分组模式
	blockMode := cipher.NewCTR(block, count)
	//执行加密、解密操作
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	//返回明文或密文
	return message
}
