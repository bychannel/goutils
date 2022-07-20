package encryutil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"strings"
)

// 不可逆加密算法

// MD5 MD5加密算法
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Md5Check 对比加密串是否和给定的明文一致
// 因为不可解密，所以只能对明文再进行一次加密，对比加密后的结果进行判定
func Md5Check(data, encrypted string) bool {
	return strings.EqualFold(MD5(data), encrypted)
}

// Sha1 SHA1加密算法
func Sha1(data string) string {
	s1 := sha1.New()
	s1.Write([]byte(data))
	return hex.EncodeToString(s1.Sum([]byte("")))
}

// Sha256 SHA256加密算法
func Sha256(data string) string {
	s256 := sha256.New()
	s256.Write([]byte(data))
	return hex.EncodeToString(s256.Sum([]byte("")))
}

// Sha512 SHA512加密算法
func Sha512(data string) string {
	s512 := sha512.New()
	s512.Write([]byte(data))
	return hex.EncodeToString(s512.Sum([]byte("")))
}

// Hmac HMAC加密算法
// 注意: 此为md5编码版本
func Hmac(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Murmur(Multiply and Rotate)
