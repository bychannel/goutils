package encryutil

import (
	"fmt"
	"testing"
)

//Md5进行签名
func TestMd5Encode(t *testing.T) {
	str := "123456789q"
	fmt.Printf("签名前-----------------%s\n", str)
	fmt.Printf("签名后-----------------%s\n", MD5(str))
}

//Md5签名验证，校验签名前和签名后是否是一个字符串
func TestMd5Check(t *testing.T) {
	str := "123456789q"
	fmt.Printf("签名前-----------------%s\n", str)
	encode := MD5(str)
	fmt.Printf("签名后-----------------%s\n", encode)
	fmt.Printf("签名校验-----------------%t\n", Md5Check(str, encode))
}
