package encryutil

import (
	"fmt"
	"testing"
)

//AES加密方法
func TestAesEncrypt(t *testing.T) {
	str := "1234567789"
	//自定义加密key
	key := "ksjdu7372hsy43l;"
	encrypt, _ := AesEncrypt(str)
	fmt.Printf("使用默认key进行加密---------->%s\n", encrypt)
	aesEncrypt, _ := AesEncrypt(str, key)
	fmt.Printf("使用自定义key进行加密---------->%s\n", aesEncrypt)
}

//AES解密方法
func TestAesDecrypt(t *testing.T) {
	str := "1234567789"
	encrypt, _ := AesEncrypt(str)
	fmt.Printf("解密前---------->%s\n", encrypt)
	decrypt, err := AesDecrypt(encrypt)
	if err != nil {
		return
	}
	fmt.Printf("解密后---------->%s\n", decrypt)
}

//DES加密方法
func TestDesEncrypt(t *testing.T) {
	str := "1234567789"
	//自定义加密key
	key := "ujugygyi"
	encrypt, _ := DesEncrypt(str)
	fmt.Printf("使用默认key进行加密---------->%s\n", encrypt)
	desEncrypt, _ := DesEncrypt(str, key)
	fmt.Printf("使用自定义key进行加密---------->%s\n", desEncrypt)
}

//DES解密方法
func TestDesDecrypt(t *testing.T) {
	str := "1234567789"
	encrypt, _ := DesEncrypt(str)
	fmt.Printf("解密前---------->%s\n", encrypt)
	decrypt, err := DesDecrypt(encrypt)
	if err != nil {
		return
	}
	fmt.Printf("解密后---------->%s\n", decrypt)
}
